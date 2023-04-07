package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"google.golang.org/protobuf/types/known/timestamppb"
	doctorpb "main.go/gunk/v1/doctor"
)
type Load struct {
	ID int
	StartAt string
	EndAt string
	WorkDays string
	Address string
	Phone string
}
type LoadManageScheduleFrom struct {
	User Load
	FormError map[string]error
	CSRFToken string
}
func (h Handler)MangeSchedule(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	fmt.Println(id)
	UId ,err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    Data := LoadManageScheduleFrom{
    	User:      Load{
    		ID:      UId,
    	},
    	CSRFToken: nosurf.Token(r),
    }
	h.ParseManageScheduleTemplate(w,Data)
}

func (h Handler) MangeSchedulePost(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	UId ,err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}

	form := LoadManageScheduleFrom{}
	user := Load{
		ID:       UId,
	}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
		return
	}
    if err := user.ValidateDoctorSchedule(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseManageScheduleTemplate(w,LoadManageScheduleFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),})
		return
	}
	u, err := h.usermgmService.DoctorDetailsList(r.Context(),&doctorpb.DoctorDetailsListRequest{
		ID: int32(UId),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ , err = h.usermgmService.DoctorScheduleRegister(r.Context(),&doctorpb.DoctorScheduleRegisterRequest{
		DoctorDetailsID: u.ID,
		StartAt:         timestamppb.New(h.StringToDate(user.StartAt)),
		EndAt:           timestamppb.New(h.StringToDate(user.EndAt)),
		WorkDays:        user.WorkDays,
		Address:         user.Address,
		Phone:           user.Phone,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/doctor/%v/home",id), http.StatusSeeOther)
}
func (u Load) ValidateDoctorSchedule() error {
	return validation.ValidateStruct(&u, validation.Field(&u.StartAt,
		validation.Required.Error("StartAt not be blank"),
	),
		validation.Field(&u.EndAt,
			validation.Required.Error("EndAt can not be blank"),
		),
		validation.Field(&u.Address,
			validation.Required.Error("Address can not be blank"),
		),
		validation.Field(&u.Phone,
			validation.Required.Error("Phone can not be blank"),
		),
		validation.Field(&u.WorkDays,
			validation.Required.Error("WorkDays can not be blank"),
		),
	)
}
func(h Handler) ParseManageScheduleTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("createDoctorSchedule.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
func (h *Handler) StringToDate(date string) time.Time {
	layout := "2006-01-02 "
	fdate, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
	}
	return fdate
}