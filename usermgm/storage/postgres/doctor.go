package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//insert into doctor_type
const registerdoctor_typeQuery = `INSERT INTO doctor_type (
	doctor_type
) values(
	:doctor_type
)RETURNING *`
func(s PostGressStorage) Registerdoctortype(u storage.Doctor_type) (*storage.Doctor_type, error){
	stmt, err := s.DB.PrepareNamed(registerdoctor_typeQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Println("error is in the query section of registration section")
		return nil, err
	}
	if u.ID == 0 {
		log.Println("error is in the query section of registration section u.ID == 0")
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

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
