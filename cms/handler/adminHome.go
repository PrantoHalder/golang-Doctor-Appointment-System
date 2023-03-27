package handler

import (
	"net/http"
)

func(h Handler) AdminHome(w http.ResponseWriter, r *http.Request) {
	h.ParseAdminHomeTemplate(w,nil)
}
func(h Handler) ParseAdminHomeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("adminHome.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
