package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//insert into users
//test case done
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
func(s PostGressStorage) Register(u storage.Register) (*storage.Register, error){
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
//register appointments
//test case done
const registerAppointmentQuery = `INSERT INTO appointment (
	userid,
	doctordetailsid,
	schduleid
) values(
	:userid,
	:doctordetailsid,
	:schduleid
)RETURNING id`
func(s PostGressStorage) RegisterAppointment(u storage.Appointment) (*storage.Appointment, error){
	stmt, err := s.DB.PrepareNamed(registerAppointmentQuery)
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

//user edit
//test case done
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
//update user
//test case done
const UpdateuserQuery = `
	UPDATE users SET
		first_name = :first_name,
		last_name = :last_name,
		email = :email
	WHERE id = :id 
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateUser(u storage.UpdateUser) (*storage.UpdateUser, error) {
	stmt, err := s.DB.PrepareNamed(UpdateuserQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//delete user
//test case done
const deleteUserbyID = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteUserByID(id int) error {
	res, err := s.DB.Exec(deleteUserbyID, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}
//user list
//test case done
const listUserQuery = `

SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
	deleted_at IS NULL
	AND 
	role = 'user'
	AND 
    (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListUser(uf storage.UserFilter) ([]storage.User, error) {
	var listUser []storage.User
	if err := s.DB.Select(&listUser, listUserQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
//edit user status

const EditUserStatusQuery = `SELECT id,is_active
FROM users
WHERE
id =$1
AND
role ='user'
AND
deleted_at IS NULL`

func (s PostGressStorage) EditUserStatus(id int) (*storage.UpdateStatus, error) {
	var listUser storage.UpdateStatus
	if err := s.DB.Get(&listUser,EditUserStatusQuery,id); err != nil {
		return nil, err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update user status
//test case done
const UpdateUserStatusQuery = `
	UPDATE users SET
		is_active = :is_active
	WHERE id = :id 
	AND
	role ='user'
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateUserStatus(u storage.UpdateStatus) (*storage.UpdateStatus, error) {
	stmt, err := s.DB.PrepareNamed(UpdateUserStatusQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//show doctor list to user
//test case done
const ShowDoctorListToUserQuery = `SELECT doctordetails.id, users.first_name, users.last_name,doctordetails.degree,doctortype.doctortype,doctordetails.gender
FROM users
FULL OUTER JOIN doctordetails ON doctordetails.userid = users.id
FULL OUTER JOIN doctortype ON doctordetails.userid = doctortype.id
WHERE doctordetails.doctortypeid = $1
`
func (s PostGressStorage) ShowDoctorListToUser(id int) ([]storage.ShowDoctorToPatient, error) {
	var listUser []storage.ShowDoctorToPatient
	if err := s.DB.Select(&listUser,ShowDoctorListToUserQuery,id); err != nil {
		return nil, err
	}
	if listUser == nil {
     return nil,fmt.Errorf("unable to find username")
	}
	return listUser, nil
}
//appointment status
const AppinmentStatusQuery = `SELECT appointment.id, users.first_name, users.last_name,appointment.is_appointed,appointment.timeslot
FROM users
FULL OUTER JOIN doctordetails ON doctordetails.userid = users.id
FULL OUTER JOIN appointment ON doctordetails.userid = appointment.doctordetailsid
WHERE appointment.userid = $1`

func (s PostGressStorage)AppinmentStatus (id int) ([]storage.AppontmentStatus, error) {
	var listUser []storage.AppontmentStatus
	if err := s.DB.Select(&listUser,AppinmentStatusQuery,id); err != nil {
		return nil, err
	}
	if listUser == nil {
     return nil,fmt.Errorf("unable to find username")
	}
	return listUser, nil
}
