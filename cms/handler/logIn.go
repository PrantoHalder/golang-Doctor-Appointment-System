package handler

import (
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type LoginUser struct {
	Username  string 
	Password  string
	Loginas   []string
	FormError map[string]error
	CSRFToken string
}

func (h Handler) Login (w http.ResponseWriter, r *http.Request){
	h.ParseLoginTemplates(w,nil)
}

func (h Handler) ParseLoginTemplates (w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		log.Fatal("can not look up login.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
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
		if err := lf.validate(); err != nil {
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


	//patient log in
	for _, value := range lf.Loginas {
		if value == "Patient" {
			
		}
	}

	

	
	h.ParseLoginTemplates(w, nil)
}
func (lu LoginUser) validate() error {
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