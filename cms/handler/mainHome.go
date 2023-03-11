package handler

import (
	"log"
	"net/http"
)

func(h Handler) MainHome(w http.ResponseWriter, r *http.Request) {
	h.ParseMainHomeTemplate(w,nil)
}
func(h Handler) ParseMainHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("mainHome.html")
	if t == nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up mainHome.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}