package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type DoctorStore interface {
	GetDoctorByUsername(string) (*storage.User, error)
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
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
		fmt.Println("the error is in the core layer in GetStatusbyUsernameCore after cu.store.GetUserByUsername(login) ")
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil{
		fmt.Println("the error is in the core layer in GetStatusbyUsernameCore after bcrypt.CompareHashAndPassword ")
		return nil, err
	}
	return user,nil
}

//doctor_type create
func (cu CoreDoctor)Registerdoctortype(u storage.Doctor_type) (*storage.Doctor_type, error){
	ru, err := cu.store.Registerdoctortype(u)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after cu.store.Register")
		return nil, err
	}
	if ru == nil {
		fmt.Println("the error is in the core layer in Register after ru == nil")
		return nil, fmt.Errorf("enable to register")
	}
	return ru, nil
}
