package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	userpb "main.go/gunk/v1/user"
)

func (h Handler) EditUserStatus(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	Id,err :=strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	res,err := h.usermgmService.EditPatientStatus(r.Context(),&userpb.EditPatientStatusRequest{
		ID: int32(Id),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if res.IsActive{
		res.IsActive=false
		_,err := h.usermgmService.UpdatePatientStatus(r.Context(),&userpb.UpdatePatientStatusRequest{
			ID:       res.ID,
			IsActive: res.IsActive,
		})
        if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
	if !res.IsActive{
		res.IsActive=true
		_,err := h.usermgmService.UpdatePatientStatus(r.Context(),&userpb.UpdatePatientStatusRequest{
			ID:       res.ID,
			IsActive: res.IsActive,
		})
        if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
	http.Redirect(w,r,"/admin/showpatient",http.StatusSeeOther)
}