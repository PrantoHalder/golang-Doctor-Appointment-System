package handler

import (
	"net/http"
	"time"
)

type DoctorSchedule struct {
	ID int
	FirstName string
	StartAt time.Time
	EndAt time.Time
	WorkDays string
	Address string
	Phone string
}
type DoctorScheduleForm struct {
	User DoctorSchedule
	FormError map[string]error
	CSRFToken string
}

func(h Handler) CreateSchedule (w http.ResponseWriter, r *http.Request){
	h.ParseDoctorScheduleTemplate(w,nil)
}
func (h Handler)ParseDoctorScheduleTemplate(w http.ResponseWriter,data any){
	t := h.Templates.Lookup("createDoctorSchedule.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}