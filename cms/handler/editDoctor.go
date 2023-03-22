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

type DoctorUpdate struct {
	ID int
	FirstName string
	LastName string
	Email string
}
type UpdateDoctorForm struct {
	User DoctorUpdate
	CSRFToken string
	FormError map[string]error
}

func (h Handler)EditDcotor(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	Id,err :=strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	res,err := h.usermgmService.EditDoctorAdmin(r.Context(),&adminpb.EditDoctorAdminRequest{
		ID: int32(Id),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	data := DoctorRegisterLoadFrom{
		User:DoctorCreate{
			ID: Id,
			FirstName:res.User.FirstName,
			LastName: res.User.LastName, 
			Email: res.User.Email,
		},
		CSRFToken:nosurf.Token(r) ,
	}
	h.ParseEditDoctorTemplate(w,data)
}
func (h Handler) ParseEditDoctorTemplate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("doctorEdit.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (h Handler) UpdateDoctor(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	form := UpdateDoctorForm{}
	user := DoctorUpdate{}
	user = DoctorUpdate{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseEditDoctorTemplate(w,UpdateDoctorForm{
			User:      user,
			CSRFToken: nosurf.Token(r),
			FormError: form.FormError,
		})
		return
	}
	
	_, err = h.usermgmService.UpdateDoctorAdmin(r.Context(),&adminpb.UpdateDoctorAdminRequest{
		ID:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w,r,"/admin/showdoctor",http.StatusSeeOther)
}

func (u DoctorUpdate) Validate() error {
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
func(h Handler)DeleteDoctor(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	_,err = h.usermgmService.DeleteDoctorByID(r.Context(),&adminpb.DeleteAdminByIDRequest{
		ID: int32(uID),
	})
	if err != nil {
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/admin/showdoctor",http.StatusSeeOther)
}