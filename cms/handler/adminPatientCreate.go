package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
	adminpb "main.go/gunk/v1/admin"
)

type PatientCreate struct {
	ID int
	FirstName string       
	LastName  string       
	Email     string
	Role      string       
	Username  string
	Password  string       
	Is_active    bool         
}

type PatientRegisterLoadFrom struct{
	User PatientCreate
    FormError map[string]error
	CSRFToken string
}
func (h Handler) PatientRegister (w http.ResponseWriter, r *http.Request){
	h.ParsePatientRegisterTemplates(w,PatientRegisterLoadFrom{
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) PatientRegisterPost (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}

	form := PatientRegisterLoadFrom{}
	user := PatientCreate{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
		return
	}

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParsePatientRegisterTemplates(w,PatientRegisterLoadFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),})
		return
	}

	h.usermgmService.RegisterPatient(r.Context(),&adminpb.RegisterPatientRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role :user.Role,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	})
	http.Redirect(w, r, fmt.Sprintln("/admin/showpatient"), http.StatusSeeOther)
}
func (h Handler) ParsePatientRegisterTemplates(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("adminPatientCreate.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func (u PatientCreate) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3, 45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Username,
			validation.Required.Error("username cannot be blank"),
			validation.Length(4, 10).Error("fast name must be between 4 to 10 characters"),
		),
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
			validation.Length(4, 10).Error("Role should Be user"),
			validation.Match(regexp.MustCompile(`^user$`)).
		    Error("Role should Be user"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
			is.Email.Error("email should be in valid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("password must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}