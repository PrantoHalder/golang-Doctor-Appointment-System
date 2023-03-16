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
func(s PostGressStorage) Register(u storage.Patient) (*storage.Patient, error){
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
//update user
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
const listUserQuery = `

SELECT id,first_name,last_name,email,is_active
FROM users
WHERE
	deleted_at IS NULL
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
