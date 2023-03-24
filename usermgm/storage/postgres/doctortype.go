package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

//insert into doctor_type
// test case done
const registerdoctor_typeQuery = `INSERT INTO doctorType (doctortype)
VALUES (:doctortype)
ON CONFLICT (doctortype) DO NOTHING
RETURNING *;`
func(s PostGressStorage) Registerdoctortype(u storage.DoctorType) (*storage.DoctorType, error){
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
//edit doctor type
// test case done
const EditDcotorTypeQuery = `SELECT id,doctorType
FROM doctortype
WHERE
id =$1
AND
deleted_at IS NULL`

func (s PostGressStorage) EditDoctorType(id int) (*storage.DoctorType, error) {
	var listUser storage.DoctorType
	if err := s.DB.Get(&listUser,EditDcotorTypeQuery,id); err != nil {
		log.Println("error is in the query section of usermgm edit Doctortype section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin Doctortype section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//update doctor type
const UpdateDoctorTypeQuery = `
	UPDATE doctorType SET
		doctortype = :doctortype
	WHERE id = :id 
	AND 
	deleted_at is NULL
	RETURNING id;
	`

func (s PostGressStorage) UpdateDoctorType(u storage.DoctorType) (*storage.DoctorType, error) {
	stmt, err := s.DB.PrepareNamed(UpdateDoctorTypeQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil

}
//delete doctor type
const deleteDoctorTypeID = `UPDATE doctorType SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL ;`

func (s PostGressStorage) DeleteDoctorTypeID(id int) error {
	res, err := s.DB.Exec(deleteDoctorTypeID,id)
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
//docotor type list
const listDoctorListQuery = `

SELECT id,doctortype
FROM doctortype
WHERE
	deleted_at IS NULL
	AND 
    (doctortype ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
`

func (s PostGressStorage) ListDoctorType(uf storage.UserFilter) ([]storage.DoctorType, error) {
	var listUser []storage.DoctorType
	if err := s.DB.Select(&listUser, listDoctorListQuery, uf.SearchTerm); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return listUser, nil
}