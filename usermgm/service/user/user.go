package user

import (
	"context"

	userpb "main.go/gunk/v1/user"
	"main.go/usermgm/storage"
)

type CoreUser interface {
	Register (storage.User) (*storage.User, error) 
	
}

type UserSvc struct {
	userpb.UnimplementedUserServiceServer
	core CoreUser
}

func NewUserSvc(cu CoreUser) *UserSvc {
	return &UserSvc{
		core: cu,
	}
}

func(us UserSvc) Register(ctx context.Context,r *userpb.RegisterRequest) (*userpb.RegisterResponse, error){
 user := storage.User{
 	FirstName: r.GetFirstName(),
 	LastName:  r.GetLastName(),
 	Email:     r.GetEmail(),
 	Username:  r.GetUsername(),
 	Password:  r.GetPassword(),
 }
 if err := user.Validate(); err != nil {
   return nil,err //TODO :: will fix after cms is done
 }
   u,err:=us.core.Register(user)
   if err != nil {
	return nil,err 
   }
   return &userpb.RegisterResponse{
   	User: &userpb.User{
   		ID:        int32(u.ID),
   		FirstName: u.FirstName,
   		LastName:  u.LastName,
   		Username:  u.Username,
   		Email:     u.Email,
   		Status:    u.Status,
   		Role:      u.Role,
   	},
   },nil
}

