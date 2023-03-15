package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//doctor login
const getDoctorByUsernameQuery=`SELECT *  
FROM users
WHERE
username = $1
AND
role ='doctor'
AND
is_active = true
AND
deleted_at IS NULL`
func (s PostGressStorage) GetDoctorByUsername(username string) (*storage.User, error) {
	var listUser storage.User
	if err := s.DB.Get(&listUser,getDoctorByUsernameQuery,username); err != nil {
		log.Println("error is in the query section of usermgm login section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin login section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
