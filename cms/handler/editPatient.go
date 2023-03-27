package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
	adminpb "main.go/gunk/v1/admin"
)

type UserUpdate struct {
	ID int
	FirstName string
	LastName string
	Email string
}
type UpdateUserForm struct {
	User UserUpdate
	CSRFToken string
	FormError map[string]error
}

func (h Handler)EditPatient(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	Id,err :=strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	res,err := h.usermgmService.EditPatient(r.Context(),&adminpb.EditPatientRequest{
		ID: int32(Id),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	data := PatientRegisterLoadFrom{
		User:PatientCreate{
			ID: Id,
			FirstName:res.User.FirstName,
			LastName: res.User.LastName, 
			Email: res.User.Email,
		},
		CSRFToken:nosurf.Token(r) ,
	}
	h.ParseEditUserTemplate(w,data)
}
func (h Handler) ParseEditUserTemplate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("adminPatientEdit.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (h Handler) UpdatePatient(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	form := UpdateUserForm{}
	user := UserUpdate{}
	user = UserUpdate{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseEditUserTemplate(w,UpdateUserForm{
			User:      user,
			CSRFToken: nosurf.Token(r),
			FormError: form.FormError,
		})
		return
	}
	
	_, err = h.usermgmService.UpdatePatient(r.Context(),&adminpb.UpdatePatientRequest{
		ID:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w,r,"/admin/showpatient",http.StatusSeeOther)
}

func (u UserUpdate) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3, 45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
			is.Email.Error("email should be in valid format"),
		),
	)
}