package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)

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
		return nil, err
	}
	if u.ID == 0 {
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

