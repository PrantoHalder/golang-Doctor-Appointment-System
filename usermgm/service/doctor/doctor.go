package doctor

import (
	"context"
	"fmt"

	doctorpb "main.go/gunk/v1/doctor"
	"main.go/usermgm/storage"
)

type CoreDoctor interface {
	GetDoctorbyUsernameCore(storage.Login) (*storage.User, error)
	RegisterDoctorCore(storage.Doctor)(*storage.Doctor,error)
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
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsActive:  u.Is_active,
			Username:  u.Username,
			Email:     u.Email,
		},
	}, nil
}
// register doctor
func (us DoctorSvc) RegisterDoctor(ctx context.Context,r *doctorpb.RegisterDoctorRequest) (*doctorpb.RegisterDoctorResponse, error){
	fmt.Println("req service", r)
	user :=storage.Doctor{
		UserID:       int(r.GetUserID()),
		DoctorTypeID: int(r.GetUserID()),
		Degree:       r.GetDegree(),
		Gender:       r.GetGender(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after RegisterDoctor")
		return nil, err
	}

	u, err := us.core.RegisterDoctorCore(user)
	fmt.Println("req response", u)
	if err != nil {
		fmt.Println("response error", err.Error())
		fmt.Println("the error is in the serveice layer in Login after RegisterDoctorCore(user)")
		return nil, err
	}

	return &doctorpb.RegisterDoctorResponse{
		User: &doctorpb.Doctor{
			ID:           int32(u.ID),
			UserID:       int32(u.UserID),
			DoctorTypeID: int32(u.DoctorTypeID),
			Degree:       u.Degree,
			Gender:       u.Gender,
		},
	}, nil
}