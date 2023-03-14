package user

import (
	"context"
	"fmt"

	userpb "main.go/gunk/v1/user"
	"main.go/usermgm/storage"
)

type CoreUser interface {
	Register(storage.User) (*storage.User, error)
	RegisterAdmin(storage.User) (*storage.User, error)
	GetUserbyUsernameCore(storage.Login) (*storage.User, error)
	GetAdminbyUsernameCore(storage.Login) (*storage.User, error)
	RegisterDoctor(storage.User) (*storage.User, error)
	GetDoctorbyUsernameCore(storage.Login) (*storage.User, error)
	EditUserCore(int) (*storage.User, error)
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
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
//user register
func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.Register(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.RegisterResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}

//admin register
func (us UserSvc) RegisterAdmin(ctx context.Context, r *userpb.RegisterAdminRequest) (*userpb.RegisterAdminResponse, error) {
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
	return &userpb.RegisterAdminResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}

//doctor register
func (us UserSvc) RegisterDoctor(ctx context.Context, r *userpb.RegisterDoctorRequest) (*userpb.RegisterDoctorResponse, error) {
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
	u, err := us.core.RegisterDoctor(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.RegisterDoctorResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
		},
	}, nil
}
//doctor type register
func (us UserSvc) RegisterDoctorType(ctx context.Context, r *userpb.RegisterDoctorTypeRequest) (*userpb.RegisterDoctorTypeResponse, error) {
	user := storage.Doctor_type{
		ID:         int(r.GetID()),
		DoctorType: r.GetDoctorType(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.Registerdoctortype(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.RegisterDoctorTypeResponse{
		User: &userpb.DoctorType{
			ID:         int32(u.ID),
			DoctorType: u.DoctorType,
		},
	}, nil
}

//user login
func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}

	u, err := us.core.GetUserbyUsernameCore(login)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.GetStatusbyUsernameCore(login)")
		return nil, err
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
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

//admin login
func (us UserSvc) AdminLogin(ctx context.Context, r *userpb.AdminLoginRequest) (*userpb.AdminLoginResponse, error) {
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

	return &userpb.AdminLoginResponse{
		User: &userpb.User{
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
//doctor login
func (us UserSvc) DoctorLogin(ctx context.Context, r *userpb.DoctorLoginRequest) (*userpb.DoctorLoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}

	u, err := us.core.GetDoctorbyUsernameCore(login)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.GetStatusbyUsernameCore(login)")
		return nil, err
	}

	return &userpb.DoctorLoginResponse{
		User: &userpb.User{
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
//edit user
func (us UserSvc) UserEdit(cxt context.Context,r *userpb.UserEditRequest) (*userpb.UserEditResponse, error){
	user := storage.Edit{
		ID: int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}
	u, err := us.core.EditUserCore(user.ID)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.EditUserCore(user.ID)")
		return nil, err
	}
	return &userpb.UserEditResponse{
		User: &userpb.User{
			ID: int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsActive:  u.Is_active,
		},
	}, nil
}