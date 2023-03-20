package handler

import (
	"log"
	"net/http"
	adminpb "main.go/gunk/v1/admin"
)



type AdminFilter struct {
	Users []AdminCreate
	SearchTerm string
}


func (h Handler) ShowAdmin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	ListUser, err := h.usermgmService.AdminList(r.Context(),&adminpb.AdminListRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := []AdminCreate{}
	if ListUser != nil {
		for _, v := range ListUser.GetUsers() {
			data = append(data,AdminCreate{
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
	Data := AdminFilter{
		Users:      data,
		SearchTerm: st,
	}

	h.ParseAdminListTemplate(w, Data)
}

func (h Handler) ParseAdminListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("listAdmin.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}