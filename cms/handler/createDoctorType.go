package handler

import (
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	doctortypepb "main.go/gunk/v1/doctortype"
)

type DoctorTypeCreate struct {
	ID int
	DoctorType string       
}
type DoctorTypeRegisterLoadFrom struct{
	User DoctorTypeCreate
    FormError map[string]error
	CSRFToken string
}
func (h Handler) DoctorTypeRegister (w http.ResponseWriter, r *http.Request){
	h.ParseDoctorTypeRegisterTemplates(w,DoctorTypeRegisterLoadFrom{
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) DoctorTypeRegisterPost (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}

	form := DoctorTypeRegisterLoadFrom{}
	user := DoctorTypeCreate{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
		return
	}

	if err := user.ValidateDoctorType(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseDoctorTypeRegisterTemplates(w,DoctorTypeRegisterLoadFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),})
		return
	}

	h.usermgmService.RegisterDoctorType(r.Context(),&doctortypepb.RegisterDoctorTypeRequest{
		DoctorType: user.DoctorType,
	})
	http.Redirect(w, r, fmt.Sprintln("/admin/home"), http.StatusSeeOther)
}
func (h Handler)ParseDoctorTypeRegisterTemplates(w http.ResponseWriter,data any){
	t := h.Templates.Lookup("doctorTypeCreate.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (u DoctorTypeCreate) ValidateDoctorType() error {
	return validation.ValidateStruct(&u, validation.Field(&u.DoctorType,
		validation.Required.Error("DoctorTypecan not be blank"),
		validation.Length(3, 45).Error("DoctorType must be between 3 to 45 characters"),
	),)
}