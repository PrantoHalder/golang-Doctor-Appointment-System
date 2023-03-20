package handler

import (
	"log"
	"net/http"
	doctorpb "main.go/gunk/v1/doctor"
)



type DoctorFilter struct {
	Users []DoctorCreate
	SearchTerm string
}


func (h Handler) ShowDoctor(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	ListUser, err := h.usermgmService.DoctorList(r.Context(),&doctorpb.DoctorListRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := []DoctorCreate{}
	if ListUser != nil {
		for _, v := range ListUser.GetUser() {
			data = append(data,DoctorCreate{
				ID:        int(v.ID),
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Email:     v.Email,
				Role:      v.Role,
				Username:  v.Username,
				Is_active:    v.IsActive,
			} )
		}
	}	
	Data := DoctorFilter{
		Users:      data,
		SearchTerm: st,
	}

	h.ParseDoctorListTemplate(w, Data)
}

func (h Handler) ParseDoctorListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("listDoctor.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}