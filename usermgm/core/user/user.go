package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	GetAdminByUsername(string) (*storage.User, error)
	RegisterAdmin(storage.User) (*storage.User, error)
	RegisterDoctor(storage.User) (*storage.User, error)
	GetDoctorByUsername(string) (*storage.User, error)
	EditUser(int) (*storage.User, error)
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
}

type CoreUser struct {
	store UserStore
}

func NewCoreUser(us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

// user registration function
func (cu CoreUser) Register(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.Register(u)
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
// user registration function
func (cu CoreUser) RegisterDoctor(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.RegisterDoctor(u)
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

// Admin registration function
func (cu CoreUser) RegisterAdmin(u storage.User) (*storage.User, error) {
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
//user login
func (cu CoreUser) GetUserbyUsernameCore(login storage.Login) (*storage.User, error){
    user,err := cu.store.GetUserByUsername(login.Username) 
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

//admin login
func (cu CoreUser) GetAdminbyUsernameCore(login storage.Login) (*storage.User, error){
    user,err := cu.store.GetAdminByUsername(login.Username) 
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
//admin login
func (cu CoreUser) GetDoctorbyUsernameCore(login storage.Login) (*storage.User, error){
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
//edit user
func (cu CoreUser) EditUserCore(id int) (*storage.User, error){
	var usr storage.User
	user ,err := cu.store.EditUser(usr.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//doctor_type create
func (cu CoreUser)Registerdoctortype(u storage.Doctor_type) (*storage.Doctor_type, error){
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
