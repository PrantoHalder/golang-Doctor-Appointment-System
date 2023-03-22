package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)
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
func(s PostGressStorage) RegisterAdmin(u storage.Admin) (*storage.Admin, error){
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
//doctor register by admin
const registerDoctorAminQuery = `INSERT INTO users (
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
func(s PostGressStorage) RegisterDoctorAdmin(u storage.DoctorU) (*storage.DoctorU, error){
	stmt, err := s.DB.PrepareNamed(registerDoctorAminQuery)
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
//user register by admin
const registerpatientByAdminQuery = `INSERT INTO users (
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
func(s PostGressStorage) RegisterPatient(u storage.User) (*storage.User, error){
	stmt, err := s.DB.PrepareNamed(registerpatientByAdminQuery)
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
// admin edit
const EditAdminQuery = `SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
id =$1
AND
role ='admin'
AND
deleted_at IS NULL`

func (s PostGressStorage) EditAdmin(id int) (*storage.Admin, error) {
	var listUser storage.Admin
	if err := s.DB.Get(&listUser,EditAdminQuery,id); err != nil {
		log.Println("error is in the query section of usermgm edit user section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin edit user section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update Admin
const UpdateAdminQuery = `
	UPDATE users SET
		first_name = :first_name,
		last_name = :last_name,
		email = :email
	WHERE id = :id 
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateAdmin(u storage.UpdateUser) (*storage.UpdateUser, error) {
	stmt, err := s.DB.PrepareNamed(UpdateAdminQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//delete admin
const deleteAdminbyID = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteAdminByID(id int) error {
	res, err := s.DB.Exec(deleteAdminbyID, id)
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
// doctor edit
const EditDoctorQuery = `SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
id =$1
AND
role ='doctor'
AND
deleted_at IS NULL`

func (s PostGressStorage) EditDoctor(id int) (*storage.DoctorU, error) {
	var listUser storage.DoctorU
	if err := s.DB.Get(&listUser,EditDoctorQuery,id); err != nil {
		log.Println("error is in the query section of usermgm edit user section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin edit user section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update doctor
const UpdateDoctorQuery = `
	UPDATE users SET
		first_name = :first_name,
		last_name = :last_name,
		email = :email
	WHERE id = :id 
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateDoctor(u storage.UpdateUser) (*storage.UpdateUser, error) {
	stmt, err := s.DB.PrepareNamed(UpdateDoctorQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//delete doctor
const deleteDoctorbyID = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteDoctorByID(id int) error {
	res, err := s.DB.Exec(deleteDoctorbyID, id)
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
//edit patient
const EditPatientQuery = `SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
id =$1
AND
role ='user'
AND
deleted_at IS NULL`

func (s PostGressStorage) EditPatient(id int) (*storage.Patient, error) {
	var listUser storage.Patient
	if err := s.DB.Get(&listUser,EditPatientQuery,id); err != nil {
		log.Println("error is in the query section of usermgm edit user section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin edit user section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update patient
const UpdatePatientQuery = `
	UPDATE users SET
		first_name = :first_name,
		last_name = :last_name,
		email = :email
	WHERE id = :id 
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdatePatient(u storage.UpdateUser) (*storage.UpdateUser, error) {
	stmt, err := s.DB.PrepareNamed(UpdatePatientQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//delete patient
const deletePatientbyID = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeletePatientByID(id int) error {
	res, err := s.DB.Exec(deletePatientbyID, id)
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
//admin list
const listAdminQuery = `

SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
	deleted_at IS NULL
	AND 
	role = 'admin'
	AND 
    (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListAdmin(uf storage.UserFilter) ([]storage.Admin, error) {
	var listUser []storage.Admin
	if err := s.DB.Select(&listUser, listAdminQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}
