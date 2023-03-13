package handler

import (
	"log"
	"net/http"
)


func (h Handler) PatientHome(w http.ResponseWriter, r *http.Request) {
  h.ParsepPatientTempate(w, nil)
}
func (h Handler) ParsepPatientTempate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("patientHome.html")
	if t == nil {
		log.Fatal("can not look up login.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up login.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
