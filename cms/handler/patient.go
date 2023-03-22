package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
	adminpb "main.go/gunk/v1/admin"
)


func (h Handler) PatientHome(w http.ResponseWriter, r *http.Request) {
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
	h.ParsepPatientTempate(w,data)
}
func (h Handler) ParsepPatientTempate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("patientHome.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
