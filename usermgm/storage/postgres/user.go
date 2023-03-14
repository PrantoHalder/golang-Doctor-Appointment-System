package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//insert into users
const registerQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password
)RETURNING *`
func(s PostGressStorage) Register(u storage.User) (*storage.User, error){
	stmt, err := s.DB.PrepareNamed(registerQuery)
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

//admin register
const registerAdminQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	role
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:role
)RETURNING *`
func(s PostGressStorage) RegisterAdmin(u storage.User) (*storage.User, error){
	stmt, err := s.DB.PrepareNamed(registerAdminQuery)
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
//doctor register
const registerdoctorQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	role
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:role
)RETURNING *`
func(s PostGressStorage) RegisterDoctor(u storage.User) (*storage.User, error){
	stmt, err := s.DB.PrepareNamed(registerdoctorQuery)
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

//login
const getUserByUsernameQuery=`SELECT *  
FROM users
WHERE
username = $1
AND
role = 'user'
AND
deleted_at IS NULL`
func (s PostGressStorage) GetUserByUsername(username string) (*storage.User, error) {
	var listUser storage.User
	if err := s.DB.Get(&listUser, getUserByUsernameQuery,username); err != nil {
		log.Println("error is in the query section of usermgm login section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 login section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}

//admin login
const getAdminByUsernameQuery=`SELECT *  
FROM users
WHERE
username = $1
AND
role ='admin'
AND
is_active = true
AND
deleted_at IS NULL`
func (s PostGressStorage) GetAdminByUsername(username string) (*storage.User, error) {
	var listUser storage.User
	if err := s.DB.Get(&listUser,getAdminByUsernameQuery,username); err != nil {
		log.Println("error is in the query section of usermgm login section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin login section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//admin login
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
//user edit
const EditUserQuery = `SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
id =$1
AND
role ='user'
AND
deleted_at IS NULL`

func (s PostGressStorage) EditUser(id int) (*storage.User, error) {
	var listUser storage.User
	if err := s.DB.Get(&listUser,EditUserQuery,id); err != nil {
		log.Println("error is in the query section of usermgm edit user section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin edit user section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
