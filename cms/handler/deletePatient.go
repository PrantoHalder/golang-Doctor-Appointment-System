package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	adminpb "main.go/gunk/v1/admin"
)


func(h Handler)DeletePatient(w http.ResponseWriter, r *http.Request){
	fmt.Println("check-1")
	id := chi.URLParam(r,"id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	_,err = h.usermgmService.DeletePatient(r.Context(),&adminpb.DeletePatientRequest{
		ID: int32(uID),
	})
	if err != nil {
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/admin/showpatient",http.StatusSeeOther)
}