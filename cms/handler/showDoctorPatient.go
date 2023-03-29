package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	userpb "main.go/gunk/v1/user"
)
type ShowDoctorToPatient struct {
	ID         int    
	FirstName  string 
	LastName   string 
	Degree     string 
	DoctorType string
	Gender     string 
}
type ShowDoctorFilter struct {
	Users []ShowDoctorToPatient
}
func (h Handler) ShowDoctorPatient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	UId ,err:= strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	ListUser, err := h.usermgmService.ShowDoctorlistPatient(r.Context(), &userpb.ShowDoctorlistPatientRequest{
		Id: int32(UId),
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
    
	data := []ShowDoctorToPatient{}
	if ListUser != nil {
		for _, v := range ListUser.GetDoctorList() {
			data = append(data, ShowDoctorToPatient{
				ID:         UId,
				FirstName:  v.FirstName,
				LastName:   v.LastName,
				Degree:     v.Degree,
				DoctorType: v.DoctorType,
				Gender:     v.Gender,
			})
		}
	}
	Data := ShowDoctorFilter{
		Users: data,
	}
	h.ParseShowDoctorToPatientTemplate(w, Data)
}
func (h Handler) ParseShowDoctorToPatientTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("showDoctorToPatient.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
