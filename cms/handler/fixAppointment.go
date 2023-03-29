package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func(h Handler) FixAppointment(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	UId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w,"internal server error", http.StatusInternalServerError)
	}
	fmt.Println(UId)

}
func (h Handler) ParseFixAppointmentTemplate(w http.ResponseWriter, data any){
	t := h.Templates.Lookup("fixAppointment.html")
	if t == nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}