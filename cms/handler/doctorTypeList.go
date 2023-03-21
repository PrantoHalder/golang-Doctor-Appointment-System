package handler

import (
	"log"
	"net/http"

	doctortypepb "main.go/gunk/v1/doctortype"
)



type DoctTypeorFilter struct {
	Users []DoctorTypeCreate
	SearchTerm string
}


func (h Handler) ShowDoctorType(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	ListDoctorType, err := h.usermgmService.DoctorTypeList(r.Context(),&doctortypepb.DoctorTypeListRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
    
	data := []DoctorTypeCreate{}
	if ListDoctorType != nil {
		for _, v := range ListDoctorType.GetDoctorType() {
			data = append(data,DoctorTypeCreate{
				ID:         int(v.ID),
				DoctorType: v.DoctorType,
			} )
		}
	}	


	Data := DoctTypeorFilter{
		Users:      data,
		SearchTerm: st,
	}
	h.ParseDoctorTypeListTemplate(w, Data)
}

func (h Handler) ParseDoctorTypeListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("patientDoctorTypeList.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
