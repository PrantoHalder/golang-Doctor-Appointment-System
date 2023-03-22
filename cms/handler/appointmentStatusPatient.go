package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	userpb "main.go/gunk/v1/user"
)

type AppontmentStatus struct {
	ID           int    `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Is_Appointed bool   `db:"is_appointed"`
	TimeSlot     string `db:"timeslot"`
}
type AppontmentStatusFilter struct {
	Users []AppontmentStatus
}

func(h Handler) AppointmentStatus(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	UId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w,"internal server error", http.StatusInternalServerError)
	}
	ListUser, err := h.usermgmService.AppoinmentStatus(r.Context(),&userpb.AppoinmentStatusRequest{
		ID: int32(UId),
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := []AppontmentStatus{}
	if ListUser != nil {
		for _, v := range ListUser.GetAppontmentStatus() {
			data = append(data,AppontmentStatus{
				ID:           UId,
				FirstName:    v.FirstName,
				LastName:     v.LastName,
				Is_Appointed: v.Is_Appointed,
				TimeSlot:     v.TimeSlot,
			} )
		}
	}	
	Data := AppontmentStatusFilter{
		Users: data,
	}

	h.ParseAppointmentStatusPatientTemplate(w, Data)
}

func (h Handler) ParseAppointmentStatusPatientTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("appointmentStatusPatient.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}