package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler)MangeSchedule(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	fmt.Println(id)
	fmt.Println("inside manage schedule ")
}