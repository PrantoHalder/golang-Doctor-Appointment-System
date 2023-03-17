package handler

import (
	"fmt"
	"log"
	"net/http"

	userpb "main.go/gunk/v1/user"
)



type UserFilter struct {
	Users *userpb.UserListResponse
	SearchTerm string
}


func (h Handler) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-------check-1---------")
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Println("-------check-2---------")
	st := r.FormValue("SearchTerm")
	fmt.Println("-------check-3---------",st)
	ListUser, err := h.usermgmService.UserList(r.Context(),&userpb.UserlistRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	fmt.Println("-------check-4---------")
	fmt.Printf("%#v",ListUser)
	Data := UserFilter{
		Users:      ListUser,
		SearchTerm: st,
	}
	fmt.Println("-------check-5---------")
	fmt.Printf("%#v",Data)
	h.ParsePatientListTemplate(w, Data)
}

func (h Handler) ParsePatientListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("listPatient.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
