package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	doctorpb "main.go/gunk/v1/doctor"
	doctortypepb "main.go/gunk/v1/doctortype"
)

type Doctor struct {
	ID        int
	DoctorTypeId int  `form:"DoctorTypeId"`
	Gender    string   `form:"Gender"`
	Degree    string   `form:"Degree"`
}
type DoctorDetailsLoadForm struct {
	Type []DoctorTypeCreate
	Users Doctor
    FormError map[string]error
	CSRFToken string
}
func (h Handler) InputDoctorDeatails(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	UId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	u, err := h.usermgmService.DoctorTypeList(r.Context(), &doctortypepb.DoctorTypeListRequest{})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	data := []DoctorTypeCreate{}
	if u != nil {
		for _, v := range u.GetDoctorType() {
			data = append(data, DoctorTypeCreate{
				ID:         int(v.ID),
				DoctorType: v.DoctorType,
			})
		}
	}
	Data := DoctorDetailsLoadForm{
		Type:      data,
		Users:     Doctor{
			ID:           UId,
		},
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	}

	h.ParseInputDoctorDeatailsTemplate(w, Data)
}

func (h Handler) InputDoctorDeatailspost(w http.ResponseWriter, r *http.Request) {
	u, err := h.usermgmService.DoctorTypeList(r.Context(), &doctortypepb.DoctorTypeListRequest{})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	data := []DoctorTypeCreate{}
	if u != nil {
		for _, v := range u.GetDoctorType() {
			data = append(data, DoctorTypeCreate{
				ID:         int(v.ID),
				DoctorType: v.DoctorType,
			})
		}
	}

	id := chi.URLParam(r, "id")
	UId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	form := DoctorDetailsLoadForm{}
	user := Doctor{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%#v",user)
	user = Doctor{
		ID:           UId,
		DoctorTypeId: user.DoctorTypeId,
		Gender:       user.Gender,
		Degree:       user.Degree,
	}
	fmt.Println("===========check-3>>>>>>>>>>>")
	if err := user.ValidateDoctor(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseInputDoctorDeatailsTemplate(w, DoctorDetailsLoadForm{
			Type:      data,
			Users:     user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
	fmt.Println("===========check-4>>>>>>>>>>>")

	h.usermgmService.RegisterDoctorDetails(r.Context(), &doctorpb.RegisterDoctorDetailsRequest{
		UserID:       int32(UId),
		DoctorTypeID: int32(user.DoctorTypeId),
		Degree:       user.Degree,
		Gender:       user.Gender,
	})
	fmt.Println("===========check-5>>>>>>>>>>>")
	http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
}
func (u Doctor) ValidateDoctor() error {
	return validation.ValidateStruct(&u, validation.Field(&u.Degree,
		validation.Required.Error("Degree not be blank"),
	),
		validation.Field(&u.Gender,
			validation.Required.Error("Gender can not be blank"),
		),
		validation.Field(&u.DoctorTypeId,
			validation.Required.Error("Doctor Type name can not be blank"),
		),
	)
}

func (h Handler) ParseInputDoctorDeatailsTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("InputDoctorDeatails.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
