package handler

import (
	"fmt"
	"log"
	"net/http"

	doctortypepb "main.go/gunk/v1/doctortype"
)



type DoctTypeorFilter struct {
	Users *doctortypepb.DoctorTypeListResponse
	SearchTerm string
}


func (h Handler) ShowDoctorType(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-------check-1---------")
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Println("-------check-2---------")
	st := r.FormValue("SearchTerm")
	fmt.Println("-------check-3---------",st)
	ListDoctorType, err := h.usermgmService.DoctorTypeList(r.Context(),&doctortypepb.DoctorTypeListRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	fmt.Println("-------check-4---------")
	fmt.Printf("%#v",ListDoctorType)
	Data := DoctTypeorFilter{
		Users:      ListDoctorType,
		SearchTerm: st,
	}
	fmt.Println("-------check-5---------")
	fmt.Printf("%#v",Data)
	h.ParseDoctorTypeListTemplate(w, Data)
}

func (h Handler) ParseDoctorTypeListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("doctorTypelist.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
