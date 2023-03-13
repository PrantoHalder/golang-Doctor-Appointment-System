package handler

import (
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
	userpb "main.go/gunk/v1/user"
)

type User struct {
	ID int
	FirstName string       
	LastName  string       
	Email     string       
	Username  string
	Password  string       
	Status    bool         
}

type RegisterLoadFrom struct{
	User User
    FormError map[string]error
	CSRFToken string
}

func (h Handler) Register (w http.ResponseWriter, r *http.Request){
	h.ParseRegisterTemplates(w,RegisterLoadFrom{
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) RegisterPost (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}

	form := RegisterLoadFrom{}
	user := User{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseRegisterTemplates(w,RegisterLoadFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
	h.usermgmService.Register(r.Context(),&userpb.RegisterRequest{
		ID:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	})
	http.Redirect(w, r, fmt.Sprintln("/login"), http.StatusSeeOther)
}
func (h Handler) ParseRegisterTemplates (w http.ResponseWriter,form RegisterLoadFrom) {
	t := h.Templates.Lookup("register.html")
	if t == nil {
		log.Fatal("can not look up register.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, form); err != nil {
		log.Fatal("can not look up register.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (u User) Validate() error {
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
		
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
			is.Email.Error("email should be in valid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}