package doctortype

import (
	"fmt"

	"main.go/usermgm/storage"
)

type DoctorTypeStore interface {
	Registerdoctortype(storage.DoctorType) (*storage.DoctorType, error)
	EditDoctorType(int) (*storage.DoctorType, error)
	UpdateDoctorType(u storage.DoctorType) (*storage.DoctorType, error)
	DeleteDoctorTypeID(id int) error
	ListDoctorType(uf storage.UserFilter) ([]storage.DoctorType, error) 
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
func (cu CoreDoctorType)Registerdoctortype(u storage.DoctorType) (*storage.DoctorType, error){
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
//edit doctor type
func (cu CoreDoctorType) EditDoctorType(u storage.Edit) (*storage.DoctorType, error){
	user ,err := cu.store.EditDoctorType(u.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//update doctor type
func (cu CoreDoctorType)UpdateDoctorType(u storage.DoctorType) (*storage.DoctorType,error) {
	user ,err := cu.store.UpdateDoctorType(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//delete doctor type
func (cu CoreDoctorType)DeleteDoctorTypeID(u storage.Edit) error{
	if err :=cu.store.DeleteDoctorTypeID(u.ID);err != nil{
		return nil
	}
	return nil
}
//list doctor type
func (cu CoreDoctorType)ListDoctorType(u storage.UserFilter)([]storage.DoctorType,error){
	user,err := cu.store.ListDoctorType(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}