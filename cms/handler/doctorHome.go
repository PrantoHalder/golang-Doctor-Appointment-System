package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)
type LoadHome struct {
	ID int
}
func(h Handler) DoctorHome(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	UId ,err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    Data := LoadHome{
    	ID: UId,
    }
	h.ParseDoctorHomeTemplate(w,Data)
}
func(h Handler) ParseDoctorHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("doctorHome.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}