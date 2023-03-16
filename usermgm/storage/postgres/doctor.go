package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//doctor login by user table
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
//register doctor into doctor table

const registerDoctorQuery = `INSERT INTO doctor (
	userid,
	doctortypeid,
	degree,
	gender
) values(
	:userid,
	:doctortypeid,
	:degree,
	:gender
)RETURNING *`
func(s PostGressStorage) RegisterDoctor(u storage.Doctor) (*storage.Doctor, error){
	stmt, err := s.DB.PrepareNamed(registerDoctorQuery)
	if err != nil {
		fmt.Println("prepared error", err.Error())
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		fmt.Println("stmt error", err.Error())
		log.Println("error is in the query section of registration section")
		return nil, err
	}
	if u.ID == 0 {
		log.Println("error is in the query section of registration section u.ID == 0")
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}

	fmt.Println("doctor res", u)
	return &u, nil
}
