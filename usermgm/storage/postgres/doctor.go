package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//register doctor details into doctor table
//test case done
const registerDoctordatailsQuery = `INSERT INTO doctordetails (
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
func(s PostGressStorage) RegisterDoctorDeatils(u storage.Doctor) (*storage.Doctor, error){
	stmt, err := s.DB.PrepareNamed(registerDoctordatailsQuery)
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
//test case done
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
//test case done
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

func (s PostGressStorage) ListDoctor(uf storage.UserFilter) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, listDoctorQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
//edit Doctor details
//test case done
const EditDoctorDetailsQuery = `SELECT id,degree,gender
FROM doctordetails
WHERE
id =$1
AND
deleted_at IS NULL`

func (s PostGressStorage) EditDoctorDetails(id int) (*storage.Doctor, error) {
	var listUser storage.Doctor
	if err := s.DB.Get(&listUser,EditDoctorDetailsQuery,id); err != nil {
		return nil, err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update doctor details status
//tast case done
const UpdateDoctorDetailsQuery = `
	UPDATE doctordetails SET
		degree = :degree,
		gender = :gender
	WHERE id = :id 
	AND
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateDoctorDetails(u storage.Doctor) (*storage.Doctor, error) {
	stmt, err := s.DB.PrepareNamed(UpdateDoctorDetailsQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//edit doctor schedule
//test case done 
const EditDoctorScheduleQuery = `SELECT id,startat,endat,workdays,address,phone
FROM doctor_schedule
WHERE
id =$1
`

func (s PostGressStorage) EditDoctorSchedule(id int) (*storage.Schedule, error) {
	var listUser storage.Schedule
	if err := s.DB.Get(&listUser,EditDoctorScheduleQuery,id); err != nil {
		return nil, err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update doctor details status
//test case done
const UpdateDoctorScheduleQuery = `
	UPDATE doctor_schedule SET
		startat = :startat,
		endat = :endat,
		workdays = :workdays,
		address = :address,
		phone = :phone
	WHERE id = :id 
	RETURNING id;
	`

func (s PostGressStorage) UpdateDoctorSchedule(u storage.Schedule) (*storage.Schedule, error) {
	stmt, err := s.DB.PrepareNamed(UpdateDoctorScheduleQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//approve patient edit
//test case done
const ApproveEditQuery = `SELECT id,is_appointed,timeslot
FROM appointment
WHERE
id =$1
`
func (s PostGressStorage) ApproveEdit(id int) (*storage.Appointment, error) {
	var listUser storage.Appointment
	if err := s.DB.Get(&listUser,ApproveEditQuery,id); err != nil {
		return nil, err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
} 
//approve patient update
//
const ApproveUpdateQuery = `
	UPDATE appointment SET
		is_appointed = :is_appointed,
		timeslot = :timeslot
	WHERE id = :id 
	RETURNING id;
	`

func (s PostGressStorage) ApproveUpdate(u storage.Appointment) (*storage.Appointment, error) {
	stmt, err := s.DB.PrepareNamed(ApproveUpdateQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}