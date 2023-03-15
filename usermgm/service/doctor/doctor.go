package doctor

import (
	"context"
	"fmt"

	doctorpb "main.go/gunk/v1/doctor"
	"main.go/usermgm/storage"
)

type CoreDoctor interface {
	GetDoctorbyUsernameCore(storage.Login) (*storage.User, error)
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
}

type DoctorSvc struct {
	doctorpb.UnimplementedDoctorServiceServer
	core CoreDoctor
}

func NewDoctorSvc(cu CoreDoctor) *DoctorSvc {
	return &DoctorSvc{
		core: cu,
	}
}

//doctor login
func (us DoctorSvc) DoctorLogin(ctx context.Context, r *doctorpb.DoctorLoginRequest) (*doctorpb.DoctorLoginResponse, error) {
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

	return &doctorpb.DoctorLoginResponse{
		User: &doctorpb.User{
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
//doctor type register
func (us DoctorSvc) RegisterDoctorType(ctx context.Context, r *doctorpb.RegisterDoctorTypeRequest) (*doctorpb.RegisterDoctorTypeResponse, error) {
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
	return &doctorpb.RegisterDoctorTypeResponse{
		User: &doctorpb.DoctorType{
			ID:         int32(u.ID),
			DoctorType: u.DoctorType,
		},
	}, nil
}