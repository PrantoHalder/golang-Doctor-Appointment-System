package doctor

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type DoctorStore interface {
	GetDoctorByUsername(string) (*storage.User, error)
	RegisterDoctorDeatils(u storage.Doctor) (*storage.Doctor, error)
	RegisterDoctorSchedule(u storage.Schedule) (*storage.Schedule, error)
	ListDoctor(uf storage.UserFilter) ([]storage.DoctorU, error)
	EditDoctorDetails(id int) (*storage.Doctor, error)
	UpdateDoctorDetails(u storage.Doctor) (*storage.Doctor, error)
	EditDoctorSchedule(id int) (*storage.Schedule, error)
	UpdateDoctorSchedule(u storage.Schedule) (*storage.Schedule, error)
}

type CoreDoctor struct {
	store DoctorStore
}

func NewCoreDoctor(us DoctorStore) *CoreDoctor {
	return &CoreDoctor{
		store: us,
	}
}
//update doctor schedule
func (cu CoreDoctor) UpdateDoctorScheduleCore(u storage.Schedule) (*storage.Schedule, error) {
	user ,err := cu.store.UpdateDoctorSchedule(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//edit doctor schedule
func (cu CoreDoctor) EditDoctorScheduleCore(us storage.Edit) (*storage.Schedule, error) {
	user ,err := cu.store.EditDoctorSchedule(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//update doctor details 
func (cu CoreDoctor) UpdateDoctorDetailsCore(u storage.Doctor) (*storage.Doctor, error) {
	user ,err := cu.store.UpdateDoctorDetails(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//edit doctor details
func (cu CoreDoctor) EditDoctorDetailsCore(us storage.Edit) (*storage.Doctor, error) {
	user ,err := cu.store.EditDoctorDetails(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
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
func (cu CoreDoctor) RegisterDoctorDetailsCore(u storage.Doctor)(*storage.Doctor,error){
	ru, err := cu.store.RegisterDoctorDeatils(u)
	if err != nil {
		return nil, err
	}
	if ru == nil {
		return nil, fmt.Errorf("enable to register")
	}
	return ru, nil
}
//register doctor schedule 
func (cu CoreDoctor) RegisterDoctorScheduleCore(u storage.Schedule)(*storage.Schedule,error){
	ru, err := cu.store.RegisterDoctorSchedule(u)
	if err != nil {
		return nil, err
	}
	if ru == nil {
		return nil, fmt.Errorf("enable to register")
	}
	return ru, nil
}
//doctor list
func (cu CoreDoctor) ListDoctorCore(u storage.UserFilter) ([]storage.DoctorU,error){
	user,err := cu.store.ListDoctor(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}