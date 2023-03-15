package admin

import (
	"context"
	"fmt"

	adminpb "main.go/gunk/v1/admin"
	"main.go/usermgm/storage"
)

type CoreAdmin interface {
	RegisterAdmin(storage.User) (*storage.User, error)
	GetAdminbyUsernameCore(storage.Login) (*storage.User, error)
}

type AdminSvc struct {
	adminpb.UnimplementedAdminServiceServer
	core CoreAdmin
}

func NewAdminSvc(cu CoreAdmin) *AdminSvc {
	return &AdminSvc{
		core: cu,
	}
}

//admin register
func (us AdminSvc) RegisterAdmin(ctx context.Context, r *adminpb.RegisterAdminRequest) (*adminpb.RegisterAdminResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
		Role:      r.GetRole(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.RegisterAdmin(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.RegisterAdminResponse{
		User: &adminpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}



//admin login
func (us AdminSvc) AdminLogin(ctx context.Context, r *adminpb.AdminLoginRequest) (*adminpb.AdminLoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}

	u, err := us.core.GetAdminbyUsernameCore(login)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.GetStatusbyUsernameCore(login)")
		return nil, err
	}

	return &adminpb.AdminLoginResponse{
		User: &adminpb.User{
			ID: int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsActive:  u.Is_active,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}
