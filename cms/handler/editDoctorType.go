package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	doctortypepb "main.go/gunk/v1/doctortype"
)

type DoctorTypeUpdate struct {
	ID int
	DoctorType string
}
type UpdateDoctorTypeForm struct {
	User DoctorTypeUpdate
	CSRFToken string
	FormError map[string]error
}

func (h Handler)EditDoctorType(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	Id,err :=strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	res,err := h.usermgmService.EditDoctorType(r.Context(),&doctortypepb.EditDoctorTypeRequest{
		Id: int32(Id),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	data := DoctorTypeRegisterLoadFrom{
		User:DoctorTypeCreate{
			ID:         Id,
			DoctorType: res.User.DoctorType,
		},
		CSRFToken:nosurf.Token(r) ,
	}
	h.ParseEditDoctorTypeTemplate(w,data)
}
func (h Handler) ParseEditDoctorTypeTemplate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("editDoctorType.html")
	if t == nil {
		log.Fatal("can not look up editDoctorType.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up editDoctorType.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (h Handler) UpdateDoctorType(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	form := UpdateDoctorTypeForm{}
	user := DoctorTypeUpdate{}
	user = DoctorTypeUpdate{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseEditDoctorTypeTemplate(w,UpdateDoctorTypeForm{
			User:      user,
			CSRFToken: nosurf.Token(r),
			FormError: form.FormError,
		})
		return
	}
	
	_, err = h.usermgmService.UpdateDoctorType(r.Context(),&doctortypepb.UpdateDoctorTypeRequest{
		ID:         int32(user.ID),
		DcotorType: user.DoctorType,
	})
	if err != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w,r,"/admin/showdoctortype",http.StatusSeeOther)
}

func (u DoctorTypeUpdate) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.DoctorType,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
	)
}
func(h Handler)DeleteDoctorType(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	_,err = h.usermgmService.DeleteDoctorType(r.Context(),&doctortypepb.DeleteDoctorTypeRequest{
		Id: int32(uID),
	})
	if err != nil {
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/admin/showdoctortype",http.StatusSeeOther)
}