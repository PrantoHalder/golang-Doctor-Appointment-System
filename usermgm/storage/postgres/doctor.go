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
		return nil, err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}

//register doctor into doctor table
const registerDoctorQuery = `INSERT INTO doctordetails (
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
//register doctor schedule into doctor schedule table
const registerDoctorScheduleQuery = `INSERT INTO doctor_schedule (
	doctorid,
	startat,
	endat,
    workdays,
	address,
	phone
) values(
	:doctorid,
	:startat,
	:endat,
    :workdays,
	:address,
	:phone
)RETURNING *`
func(s PostGressStorage) RegisterDoctorSchedule(u storage.Schedule) (*storage.Schedule, error){
	stmt, err := s.DB.PrepareNamed(registerDoctorScheduleQuery)
	if err != nil {
		fmt.Println("prepared error", err.Error())
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		fmt.Println("stmt error", err.Error())
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}

	return &u, nil
}
//list doctor
const listDoctorQuery = `

SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
	deleted_at IS NULL
	AND 
	role = 'doctor'
	AND 
    (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListDoctor(uf storage.UserFilter) ([]storage.DoctorU, error) {
	var listUser []storage.DoctorU
	if err := s.DB.Select(&listUser, listDoctorQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
