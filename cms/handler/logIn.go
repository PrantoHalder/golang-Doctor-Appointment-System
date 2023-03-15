package handler

import (
	"log"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	adminpb "main.go/gunk/v1/admin"
	doctorpb "main.go/gunk/v1/doctor"
	userpb "main.go/gunk/v1/user"
)

type LoginUser struct {
	Username  string 
	Password  string
	Loginas   []string
	FormError map[string]error
	CSRFToken string
}

func (h Handler) Login (w http.ResponseWriter, r *http.Request){
	h.ParseLoginTemplates(w,LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) ParseLoginTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		log.Fatal("can not look up login.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up login.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}

func (h Handler) LoginPost (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var lf LoginUser
	if err := h.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if lf.Loginas == nil {
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
	
	for _, value := range lf.Loginas {
		if value == "Patient" {
			u,err := h.usermgmService.Login(r.Context(),&userpb.LoginRequest{
				Username: lf.Username,
				Password: lf.Password,
			})
			if err != nil {
				log.Println("the error is in the login section of cms in patient after h.usermgmService.Login")
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	        http.Redirect(w, r, "/patients/home", http.StatusSeeOther)
		}
	}
	
	for _, value := range lf.Loginas {
		if value == "Admin" {
			u,err := h.usermgmService.AdminLogin(r.Context(),&adminpb.AdminLoginRequest{
				Username: lf.Username,
				Password: lf.Password,
			})
			if err != nil {
				log.Println("the error is in the login section of cms in patient after h.usermgmService.Login")
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	        http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
		}
	}

	for _, value := range lf.Loginas {
		if value == "Doctor" {
			u,err := h.usermgmService.DoctorLogin(r.Context(),&doctorpb.DoctorLoginRequest{
				Username: lf.Username,
				Password: lf.Password,
			})
			if err != nil {
				log.Println("the error is in the login section of cms in patient after h.usermgmService.Login")
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	        http.Redirect(w, r, "/doctor/home", http.StatusSeeOther)
		}
	}	
	h.ParseLoginTemplates(w, nil)
}
func (lu LoginUser) Validate() error {
	return validation.ValidateStruct(&lu, validation.Field(&lu.Username,
		validation.Required.Error("username can not be blank"),
	),
		validation.Field(&lu.Password,
			validation.Required.Error("password can not be blank"),
		),
		validation.Field(&lu.Loginas,
			validation.Required.Error("login role can not be blank"),
		),
	)
}
func (h Handler) LogoutPatienthandler (w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func (h Handler) LogoutDoctorhandler (w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func (h Handler) LogoutAdminhandler (w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}