package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	EditUser(int) (*storage.User, error)
	RegisterPatient(storage.User) (*storage.User, error)
	UpdateUser(storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteUserByID(int) error
	ListUser(storage.UserFilter) ([]storage.User, error)
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
//edit user
func (cu CoreUser) EditUserCore(us storage.Edit) (*storage.User, error){
	user ,err := cu.store.EditUser(us.ID)
	if err != nil {
		return nil,err
	}
	if user == nil{
      return nil,err
	}
	return user,nil
}
//register user by admin
func (cu CoreUser) RegisterPatient(u storage.User) (*storage.User, error) {
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
//update user
func (cu CoreUser) UpdatePatient(u storage.UpdateUser) (*storage.UpdateUser, error) {
	user ,err := cu.store.UpdateUser(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}
//delete user
func (cu CoreUser)DeleteUserByID(u storage.Edit) error{
	if err :=cu.store.DeleteUserByID(u.ID);err != nil{
		return nil
	}
	return nil
} 
//list user
func (cu CoreUser)ListUser(u storage.UserFilter) ([]storage.User,error){
	user,err := cu.store.ListUser(u)
	if err != nil {
		return nil,err
	}
	if user == nil{
		return nil,err
	}
	return user,nil
}