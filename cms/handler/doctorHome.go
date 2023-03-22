package handler

import (
	"net/http"
)

func(h Handler) DoctorHome(w http.ResponseWriter, r *http.Request) {
	h.ParseDoctorHomeTemplate(w,nil)
}
func(h Handler) ParseDoctorHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("doctorHome.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}