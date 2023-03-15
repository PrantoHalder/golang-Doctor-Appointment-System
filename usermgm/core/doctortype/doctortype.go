package doctortype

import (
	"fmt"

	"main.go/usermgm/storage"
)

type DoctorTypeStore interface {
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
}

type CoreDoctorType struct {
	store DoctorTypeStore
}

func NewCoreDoctorType(us DoctorTypeStore) *CoreDoctorType {
	return &CoreDoctorType{
		store: us,
	}
}
//doctor_type create
func (cu CoreDoctorType)Registerdoctortype(u storage.Doctor_type) (*storage.Doctor_type, error){
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
