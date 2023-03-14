package handler

import (
	"log"
	"net/http"
)

func(h Handler) AdminHome(w http.ResponseWriter, r *http.Request) {
	h.ParseAdminHomeTemplate(w,nil)
}
func(h Handler) ParseAdminHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("adminHome.html")
	if t == nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
