package handler

import (
	"log"
	"net/http"
)

func(h Handler) DoctorHome(w http.ResponseWriter, r *http.Request) {
	h.ParseDoctorHomeTemplate(w,nil)
}
func(h Handler) ParseDoctorHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("doctorHome.html")
	if t == nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}