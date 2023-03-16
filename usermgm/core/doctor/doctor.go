package doctor

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type DoctorStore interface {
	GetDoctorByUsername(string) (*storage.User, error)
	RegisterDoctor(u storage.Doctor) (*storage.Doctor, error)
}

type CoreDoctor struct {
	store DoctorStore
}

func NewCoreDoctor(us DoctorStore) *CoreDoctor {
	return &CoreDoctor{
		store: us,
	}
}

//Doctor login
func (cu CoreDoctor) GetDoctorbyUsernameCore(login storage.Login) (*storage.User, error){
    user,err := cu.store.GetDoctorByUsername(login.Username) 
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil{
		return nil, err
	}
	return user,nil
}
//register doctor
func (cu CoreDoctor) RegisterDoctorCore(u storage.Doctor)(*storage.Doctor,error){
	ru, err := cu.store.RegisterDoctor(u)
	if err != nil {
		return nil, err
	}
	if ru == nil {
		return nil, fmt.Errorf("enable to register")
	}
	return ru, nil
}
