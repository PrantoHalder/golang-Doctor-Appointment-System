package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	doctorpb "main.go/gunk/v1/doctor"
)

func(h Handler)AppointmentList(w http.ResponseWriter, r *http.Request){
	id :=chi.URLParam(r,"id")
	UId,err := strconv.Atoi(id)
	if err != nil {
		http.Error(w,"internal server error", http.StatusInternalServerError)
	}
	u, err := h.usermgmService.DoctorDetailsEdit(r.Context(),&doctorpb.DoctorDetailsEditRequest{
		ID: int32(UId),
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	fmt.Println(u)
}