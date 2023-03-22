package login

import (
	"context"
	"fmt"

	loginpb "main.go/gunk/v1/login"
	"main.go/usermgm/storage"
)

type CoreLogin interface {
	LoginCore(login storage.Login) (*storage.User, error)
}

type LoginSvc struct {
	loginpb.UnimplementedLoginServiceServer
	core CoreLogin
}

func NewLoginSvc(cu CoreLogin) *LoginSvc {
	return &LoginSvc{
		core: cu,
	}
}

// user login
func (us LoginSvc) Login(ctx context.Context,r *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}

	u, err := us.core.LoginCore(login)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.GetStatusbyUsernameCore(login)")
		return nil, err
	}

	return &loginpb.LoginResponse{
		User: &loginpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsActive:  u.Is_active,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}
