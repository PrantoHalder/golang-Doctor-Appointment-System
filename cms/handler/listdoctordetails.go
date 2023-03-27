package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	doctorpb "main.go/gunk/v1/doctor"
)
type DoctorDetails struct {
	ID int 
	FirstName string
	LastName string
	DoctorType string
	Degree string
	Gender string
}
func (h Handler) ListDoctorDetails(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	UId ,err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	u,err := h.usermgmService.DoctorDetailsList(r.Context(),&doctorpb.DoctorDetailsListRequest{
		ID: int32(UId),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	Data := DoctorDetails{
		ID:         int(u.ID),
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		DoctorType: u.DoctorType,
		Degree:     u.Degree,
		Gender:     u.Gender,
	}
    h.ParseShowdoctordetailaTemplate(w,Data)
}
func (h Handler)ParseShowdoctordetailaTemplate(w http.ResponseWriter,data DoctorDetails){
	t := h.Templates.Lookup("listdoctordetails.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}