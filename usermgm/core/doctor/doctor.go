package doctor

import (
	"fmt"

	"main.go/usermgm/storage"
)

type DoctorStore interface {
	RegisterDoctorDeatils(u storage.DoctorDetails) (*storage.DoctorDetails, error)
	RegisterDoctorSchedule(u storage.Schedule) (*storage.Schedule, error)
	ListDoctor(uf storage.UserFilter) ([]storage.User, error)
	EditDoctorDetails(id int) (*storage.DoctorDetails, error)
	UpdateDoctorDetails(u storage.DoctorDetails) (*storage.DoctorDetails, error)
	EditDoctorSchedule(id int) (*storage.Schedule, error)
	UpdateDoctorSchedule(u storage.Schedule) (*storage.Schedule, error)
    ApproveEdit(id int) (*storage.Appointment, error)
	ApproveUpdate(u storage.Appointment) (*storage.Appointment, error)
	ListDoctorDetails(id int) (*storage.DoctorDetailsList, error)
	DoctorScheduleList(id int) ([]storage.Schedule, error)
	EditDoctorStatus(id int) (*storage.UpdateStatus, error)
	UpdateDoctorStatus(u storage.UpdateStatus) (*storage.UpdateStatus, error)
}

type CoreDoctor struct {
	store DoctorStore
}

func NewCoreDoctor(us DoctorStore) *CoreDoctor {
	return &CoreDoctor{
		store: us,
	}
}
//update docotor status
func(cu CoreDoctor)UpdateDoctorStatusCore(u storage.UpdateStatus) (*storage.UpdateStatus, error) {
	user ,err := cu.store.UpdateDoctorStatus(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//edit doctor status
func(cu CoreDoctor)EditDoctorStatusCore(u storage.Edit) (*storage.UpdateStatus, error){
	user ,err := cu.store.EditDoctorStatus(u.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//doctor schedule list
func (cu CoreDoctor)DoctorScheduleListCore(u storage.Edit) ([]storage.Schedule, error){
	user ,err := cu.store.DoctorScheduleList(u.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//doctor details list
func (cu CoreDoctor) ListDoctorDetailsCore(u storage.Edit) (*storage.DoctorDetailsList, error) {
	user ,err := cu.store.ListDoctorDetails(u.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//Approve update
func (cu CoreDoctor) ApproveUpdateCore(u storage.Appointment) (*storage.Appointment, error) {
	user ,err := cu.store.ApproveUpdate(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
// Approve edit
func (cu CoreDoctor) ApproveEditCore(us storage.Edit) (*storage.Appointment, error) {
	user ,err := cu.store.ApproveEdit(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
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
func (cu CoreDoctor) UpdateDoctorDetailsCore(u storage.DoctorDetails) (*storage.DoctorDetails, error) {
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
func (cu CoreDoctor) EditDoctorDetailsCore(us storage.Edit) (*storage.DoctorDetails, error) {
	user ,err := cu.store.EditDoctorDetails(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//register doctor
func (cu CoreDoctor) RegisterDoctorDetailsCore(u storage.DoctorDetails)(*storage.DoctorDetails,error){
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
func (cu CoreDoctor) ListDoctorCore(u storage.UserFilter) ([]storage.User,error){
	user,err := cu.store.ListDoctor(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}