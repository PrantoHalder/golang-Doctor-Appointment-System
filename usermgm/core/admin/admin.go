package admin

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type AdminStore interface {
	RegisterAdmin(u storage.User) (*storage.User, error)
	RegisterDoctorAdmin(u storage.User) (*storage.User, error)
	RegisterPatient(storage.User) (*storage.User, error)
	EditAdmin(id int) (*storage.User, error)
	UpdateAdmin(storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteAdminByID(id int) error
	EditDoctor(id int) (*storage.User, error)
	UpdateDoctor(u storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteDoctorByID(id int) error 
	EditPatient(id int) (*storage.User, error)
	UpdatePatient(u storage.UpdateUser) (*storage.UpdateUser, error)
	DeletePatientByID(id int) error
	ListAdmin(uf storage.UserFilter) ([]storage.User, error)
}

type CoreAdmin struct {
	store AdminStore
}

func NewCoreAdmin(us AdminStore) *CoreAdmin {
	return &CoreAdmin{
		store: us,
	}
}


// Admin registration function
func (cu CoreAdmin) RegisterAdmin(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.RegisterAdmin(u)
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
//doctor register by admin
func (cu CoreAdmin) RegisterDoctorAdminCore(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.RegisterDoctorAdmin(u)
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
//register user by admin
func (cu CoreAdmin) RegisterPatient(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.RegisterPatient(u)
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
//admin edit
func (cu CoreAdmin) EditAdminCore(us storage.Edit) (*storage.User, error) {
	user ,err := cu.store.EditAdmin(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//update admin
func (cu CoreAdmin) UpdateAdminCore(u storage.UpdateUser) (*storage.UpdateUser, error) {
	user ,err := cu.store.UpdateAdmin(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//delete admin
func (cu CoreAdmin) DeleteAdminByIDCore(u storage.Edit) error{
	if err :=cu.store.DeleteAdminByID(u.ID);err != nil{
		return nil
	}
	return nil
}
//Doctor edit
func (cu CoreAdmin) EditDoctorCore(us storage.Edit) (*storage.User, error) {
	user ,err := cu.store.EditDoctor(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//update Doctor
func (cu CoreAdmin) UpdateDoctorCore(u storage.UpdateUser) (*storage.UpdateUser, error) {
	user ,err := cu.store.UpdateDoctor(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//delete Doctor
func (cu CoreAdmin)DeleteDoctorByIDCore(u storage.Edit) error{
	if err :=cu.store.DeleteDoctorByID(u.ID);err != nil{
		return nil
	}
	return nil
}
//patient edit
func (cu CoreAdmin) EditPatientCore(us storage.Edit) (*storage.User, error) {
	user ,err := cu.store.EditPatient(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//update patient
func (cu CoreAdmin) UpdatePatientCore(u storage.UpdateUser) (*storage.UpdateUser, error) {
	user ,err := cu.store.UpdatePatient(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//delete Patient
func (cu CoreAdmin)DeletePatientByIDCore(u storage.Edit) error{
	if err :=cu.store.DeletePatientByID(u.ID);err != nil{
		return nil
	}
	return nil
}
// list admin
func (cu CoreAdmin)ListAdminCore(u storage.UserFilter) ([]storage.User,error){
	user,err := cu.store.ListAdmin(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}