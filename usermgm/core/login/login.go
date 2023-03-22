package login

import (
	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type LoginStore interface {
	Login(username string) (*storage.User, error)
}

type CoreLogin struct {
	store LoginStore
}

func NewCorelogin(us LoginStore) *CoreLogin {
	return &CoreLogin{
		store: us,
	}
}

//user login
func (cu CoreLogin) LoginCore(login storage.Login) (*storage.User, error){
    user,err := cu.store.Login(login.Username) 
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil{
		return nil, err
	}
	return user,nil
}