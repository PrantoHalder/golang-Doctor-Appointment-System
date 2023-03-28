package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	doctorpb "main.go/gunk/v1/doctor"
)

type Schedule struct {
	ID              int
	DoctorDetailsID int
	StartAt         time.Time
	EndAt           time.Time
	WorkDays        string
	Address         string
	Phone           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
type DoctorScheduleFilter struct {
	Users []Schedule
}

func (h Handler) ListSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	UId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := h.usermgmService.DoctorDetailsList(r.Context(), &doctorpb.DoctorDetailsListRequest{
		ID: int32(UId),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ListSchedule, err := h.usermgmService.DoctorScheduleList(r.Context(), &doctorpb.DoctorScheduleListRequest{
		ID: user.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := []Schedule{}
	if ListSchedule != nil {
		for _, v := range ListSchedule.GetSchedule() {
			data = append(data,Schedule{
				ID:              int(v.ID),
				DoctorDetailsID: int(v.DoctorDetailsID),
				StartAt:         v.StartAt.AsTime(),
				EndAt:           v.GetEndAt().AsTime(),
				WorkDays:        v.WorkDays,
				Address:         v.Address,
				Phone:           v.Phone,
			} )
		}
	}	
	Data := DoctorScheduleFilter{
		Users:      data,
	}

	h.ParseListScheduleTemplate(w, Data)
}
func (h Handler) ParseListScheduleTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("listscheduleDoctor.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
