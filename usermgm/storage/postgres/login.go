package postgres

import (
	"fmt"

	"main.go/usermgm/storage"
)

//login
const getUserByUsernameQuery=`SELECT *  
FROM users
WHERE
username = $1
AND
deleted_at IS NULL`
func (s PostGressStorage) Login(username string) (*storage.User, error) {
	var listUser storage.User
	if err := s.DB.Get(&listUser, getUserByUsernameQuery,username); err != nil {
		if err.Error() == storage.NotFound{
			return nil, fmt.Errorf(NotFound)
		}
		return nil,err
	}
	if listUser.ID == 0 {
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}