package doctor

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
	doctorpb "main.go/gunk/v1/doctor"
	"main.go/usermgm/storage"
)

type CoreDoctor interface {
	GetDoctorbyUsernameCore(storage.Login) (*storage.User, error)
	RegisterDoctorCore(storage.Doctor) (*storage.Doctor, error)
	RegisterDoctorScheduleCore(u storage.Schedule) (*storage.Schedule, error)
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

// doctor login
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
func (us DoctorSvc) RegisterDoctor(ctx context.Context, r *doctorpb.RegisterDoctorRequest) (*doctorpb.RegisterDoctorResponse, error) {
	fmt.Println("req service", r)
	user := storage.Doctor{
		UserID:       int(r.GetUserID()),
		DoctorTypeID: int(r.GetUserID()),
		Degree:       r.GetDegree(),
		Gender:       r.GetGender(),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.RegisterDoctorCore(user)
	if err != nil {
		fmt.Println("response error", err.Error())
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

// register doctor schedule
func (us DoctorSvc) DoctorScheduleRegister(ctx context.Context, r *doctorpb.DoctorScheduleRegisterRequest) (*doctorpb.DoctorScheduleRegisterResponse, error) {
	fmt.Println("req service", r)
	workday, err := json.Marshal(r.WorkDays)
	if err != nil {
		return nil, err
	}

	dbPrm := storage.Schedule{
		ID:              0,
		DoctorDetailsID: int(r.GetDoctorDetailsID()),
		StartAt:         r.GetStartAt().AsTime(),
		EndAt:           r.GetEndAt().AsTime(),
		WorkDays:        string(workday),
		Address:         r.GetAddress(),
		Phone:           r.Phone,
	}
	u, err := us.core.RegisterDoctorScheduleCore(dbPrm)
	if err != nil {
		fmt.Println("response error", err.Error())
		return nil, err
	}

	return &doctorpb.DoctorScheduleRegisterResponse{
		Schedule: &doctorpb.Schedule{
			ID:              int32(u.ID),
			DoctorDetailsID: int32(u.DoctorDetailsID),
			StartAt:         timestamppb.New(u.StartAt),
			EndAt:           timestamppb.New(u.EndAt),
			WorkDays:        u.WorkDays,
			Address:         u.Address,
			Phone:           u.Phone,
		},
	}, nil
}
