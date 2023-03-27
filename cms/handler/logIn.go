package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	loginpb "main.go/gunk/v1/login"
)
const (
	NotFound = "sql: no rows in result set"
)
type LoginUser struct {
	Username  string `form:"Username"`
	Password  string `form:"Password"`
	FormError map[string]error
	CSRFToken string
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.ParseLoginTemplates(w, LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var lf LoginUser
	if err := h.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if lf.Username == "" {
		if err := lf.Validate(); err != nil {
			if vErr, ok := err.(validation.Errors); ok {
				lf.FormError = vErr
			}
			h.ParseLoginTemplates(w, LoginUser{
				Username:  lf.Username,
				Password:  "",
				FormError: lf.FormError,
				CSRFToken: nosurf.Token(r),
			})
			return
		}
	}
	if lf.Password == "" {
		if err := lf.Validate(); err != nil {
			if vErr, ok := err.(validation.Errors); ok {
				lf.FormError = vErr
			}
			h.ParseLoginTemplates(w, LoginUser{
				Username:  "",
				Password:  lf.Password,
				FormError: lf.FormError,
				CSRFToken: nosurf.Token(r),
			})
			return
		}
	}
	u, err := h.usermgmService.Login(r.Context(), &loginpb.LoginRequest{
		Username: lf.Username,
		Password: lf.Password,
	})
	if err != nil {
        if err.Error() == NotFound  {
			formErr := make(map[string]error)
			formErr["Username"] = fmt.Errorf("credentials does not match")
			lf.FormError = formErr
			lf.CSRFToken = nosurf.Token(r)
			lf.Password = ""
			h.ParseLoginTemplates(w, lf)
			return
		}
		http.Redirect(w, r, "/inactive", http.StatusSeeOther)
	}
	if u.User.Role == "admin" {
		if u.User.IsActive {
			if err := lf.Validate(); err != nil {
				if vErr, ok := err.(validation.Errors); ok {
					lf.FormError = vErr
				}
				h.ParseLoginTemplates(w, LoginUser{
					Username:  lf.Username,
					Password:  "",
					FormError: lf.FormError,
					CSRFToken: nosurf.Token(r),
				})
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
			http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
		} else {
			h.ParseInactiveTemplates(w, nil)
		}
	}
	if u.User.Role == "doctor" {
		if u.User.IsActive {
			if err := lf.Validate(); err != nil {
				if vErr, ok := err.(validation.Errors); ok {
					lf.FormError = vErr
				}
				h.ParseLoginTemplates(w, LoginUser{
					Username:  lf.Username,
					Password:  "",
					FormError: lf.FormError,
					CSRFToken: nosurf.Token(r),
				})
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	        http.Redirect(w, r,fmt.Sprintf("/doctor/%v/home",u.User.ID), http.StatusSeeOther)
		} else {
			h.ParseInactiveTemplates(w, nil)
		}
	}
	if u.User.Role == "user" {
		if u.User.IsActive {
			if err := lf.Validate(); err != nil {
				if vErr, ok := err.(validation.Errors); ok {
					lf.FormError = vErr
				}
				h.ParseLoginTemplates(w, LoginUser{
					Username:  lf.Username,
					Password:  "",
					FormError: lf.FormError,
					CSRFToken: nosurf.Token(r),
				})
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	        http.Redirect(w, r, fmt.Sprintf("/patients/%v/home",u.User.ID), http.StatusSeeOther)
		} else {
			h.ParseInactiveTemplates(w, nil)
		}
	}
}

func (lu LoginUser) Validate() error {
	return validation.ValidateStruct(&lu, validation.Field(&lu.Username,
		validation.Required.Error("username can not be blank"),
	),
		validation.Field(&lu.Password,
			validation.Required.Error("password can not be blank"),
		),
	)
}
func (h Handler) Inactive(w http.ResponseWriter, r *http.Request) {
	h.ParseInactiveTemplates(w, nil)
}
func (h Handler) InternalServerError(w http.ResponseWriter, r *http.Request) {
	h.ParseInactiveTemplates(w, nil)
}
func (h Handler) ParseLoginTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		log.Fatal("can not look up login.html template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up login.html template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h Handler) ParseInactiveTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("Inactive.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func (h Handler) ParseInternalServerErrorTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("internalServerError.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h Handler) LogoutPatienthandler(w http.ResponseWriter, r *http.Request) {
	if err := h.sessionManager.Destroy(r.Context()); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (h Handler) LogoutDoctorhandler(w http.ResponseWriter, r *http.Request) {
	if err := h.sessionManager.Destroy(r.Context()); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (h Handler) LogoutAdminhandler(w http.ResponseWriter, r *http.Request) {
	if err := h.sessionManager.Destroy(r.Context()); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
