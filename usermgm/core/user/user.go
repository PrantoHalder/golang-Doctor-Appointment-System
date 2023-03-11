package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
}

type CoreUser struct {
	store UserStore
}


func NewCoreUser (us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

func(cu CoreUser) Register (u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashPass)
	ru ,err := cu.store.Register(u)
	if err != nil {
		return nil, err
	}
	if ru == nil{
		return nil,fmt.Errorf("enable to register")
	}
	return ru, nil
}