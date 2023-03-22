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
	h.ParseLoginTemplates(w,LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}


func (h Handler) LoginPost (w http.ResponseWriter, r *http.Request){
	
	
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
func (h Handler) ParseLoginTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}