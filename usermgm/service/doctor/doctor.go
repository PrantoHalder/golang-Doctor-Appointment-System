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
	RegisterDoctorDetailsCore(u storage.DoctorDetails)(*storage.DoctorDetails,error)
	RegisterDoctorScheduleCore(u storage.Schedule) (*storage.Schedule, error)
	ListDoctorCore(u storage.UserFilter) ([]storage.User,error)
	EditDoctorDetailsCore(us storage.Edit) (*storage.DoctorDetails, error)
	UpdateDoctorDetailsCore(u storage.DoctorDetails) (*storage.DoctorDetails, error)
	EditDoctorScheduleCore(us storage.Edit) (*storage.Schedule, error)
	UpdateDoctorScheduleCore(u storage.Schedule) (*storage.Schedule, error)
	ApproveEditCore(us storage.Edit) (*storage.Appointment, error)
	ApproveUpdateCore(u storage.Appointment) (*storage.Appointment, error)
	ListDoctorDetailsCore(u storage.Edit) (*storage.DoctorDetailsList, error)
	DoctorScheduleListCore(u storage.Edit) ([]storage.Schedule, error)
	EditDoctorStatusCore(u storage.Edit) (*storage.UpdateStatus, error)
	UpdateDoctorStatusCore(u storage.UpdateStatus) (*storage.UpdateStatus, error)
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
//update doctor status
func (us DoctorSvc)UpdateDoctorStatus(ctx context.Context,r *doctorpb.UpdateDoctorStatusRequest) (*doctorpb.UpdateDoctorStatusResponse, error){
	user := storage.UpdateStatus{
		ID:        int(r.GetID()),
		Is_active: r.GetIsActive(),
	}
	u, err := us.core.UpdateDoctorStatusCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.UpdateDoctorStatusResponse{
		IsActive:u.Is_active}, nil
}
//edit doctor status
func (us DoctorSvc) EditDoctorStatus(ctx context.Context,r *doctorpb.EditDoctorStatusRequest) (*doctorpb.EditDoctorStatusResponse, error){
	user := storage.Edit{
		ID:              int(r.GetID()),
	}
	
	u, err := us.core.EditDoctorStatusCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.EditDoctorStatusResponse{
		ID:       int32(u.ID),
		IsActive: u.Is_active,
	},nil
}
//doctor schedule list
func (us DoctorSvc) DoctorScheduleList(ctx context.Context,r *doctorpb.DoctorScheduleListRequest) (*doctorpb.DoctorScheduleListResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	u,err := us.core.DoctorScheduleListCore(user)
	if err != nil {
		return nil,err
	}
	var totalusers []*doctorpb.Schedule
	for _,value := range u {
		user:=&doctorpb.Schedule{
			ID:              int32(value.ID),
			DoctorDetailsID: int32(value.DoctorDetailsID),
			StartAt:         timestamppb.New(value.StartAt),
			EndAt:           timestamppb.New(value.EndAt),
			WorkDays:        value.WorkDays,
			Address:         value.Address,
			Phone:           value.Phone,
		}
		totalusers = append(totalusers,user)
	}
	return &doctorpb.DoctorScheduleListResponse{
		Schedule: totalusers,
	}, nil
}
//doctor details list
func (us DoctorSvc)DoctorDetailsList(ctx context.Context,r *doctorpb.DoctorDetailsListRequest) (*doctorpb.DoctorDetailsListResponse, error){
	user := storage.Edit{
		ID:              int(r.GetID()),
	}
	
	u, err := us.core.ListDoctorDetailsCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.DoctorDetailsListResponse{
		ID:         int32(u.ID),
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		DoctorType: u.DoctorType,
		Degree:     u.Degree,
		Gender:     u.Gender,
	},nil
}
func (us DoctorSvc)ApproveAppointmentUpdate(ctx context.Context,r *doctorpb.ApproveAppointmentUpdateRequest) (*doctorpb.ApproveAppointmentUpdateResponse, error){
	user := storage.Appointment{
		ID:              int(r.GetID()),
		Is_Appointed:    r.GetIs_Appointed(),
		TimeSlot:        r.GetTimeSlot(),
	}
	
	u, err := us.core.ApproveUpdateCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.ApproveAppointmentUpdateResponse{
		Is_Appointed: u.Is_Appointed,
		TimeSlot:     u.TimeSlot,
	},nil
}
//approve edit
func (us DoctorSvc)ApproveAppointmentEdit(ctx context.Context,r *doctorpb.ApproveAppointmentEditRequest) (*doctorpb.ApproveAppointmentEditResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.ApproveEditCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.ApproveAppointmentEditResponse{
		ID:           int32(u.ID),
		Is_Appointed: u.Is_Appointed,
		TimeSlot:     u.TimeSlot,
	},nil
}
//update doctor schedule
func (us DoctorSvc)DoctorScheduleUpdate(ctx context.Context,r *doctorpb.DoctorScheduleUpdateRequest) (*doctorpb.DoctorScheduleUpdateResponse, error){
	workday, err := json.Marshal(r.WorkDays)
	if err != nil {
		return nil, err
	}
	user := storage.Schedule{
		ID:              int(r.GetID()),
		StartAt:         r.GetStartAt().AsTime(),
		EndAt:           r.GetEndAt().AsTime(),
		WorkDays:        string(workday),
		Address:         r.GetAddress(),
		Phone:           r.GetPhone(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.UpdateDoctorScheduleCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.DoctorScheduleUpdateResponse{
		StartAt:  timestamppb.New(u.StartAt),
		EndAt:    timestamppb.New(u.EndAt),
		WorkDays: u.WorkDays,
		Address:  u.Address,
		Phone:    u.Phone,
	},nil
}
//edit doctor schedule 
func (us DoctorSvc) DoctorScheduleEdit(ctx context.Context,r *doctorpb.DoctorScheduleEditRequest) (*doctorpb.DoctorScheduleEditResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditDoctorScheduleCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.DoctorScheduleEditResponse{
		ID:       int32(u.ID),
		StartAt:  timestamppb.New(u.StartAt),
		EndAt:    timestamppb.New(u.EndAt),
		WorkDays: u.WorkDays,
		Address:  u.Address,
	},nil
}
//doctor details update
func (us DoctorSvc)DoctorDetailsUpdate(ctx context.Context,r *doctorpb.DoctorDetailsUpdateRequest) (*doctorpb.DoctorDetailsUpdateResponse, error){
	user := storage.DoctorDetails{
		ID:           int(r.GetID()),
		Degree:       r.GetDegree(),
		Gender:       r.GetGender(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.UpdateDoctorDetailsCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.DoctorDetailsUpdateResponse{
		Degree: u.Degree,
		Gender: u.Gender,
	},nil
}
//doctor details edit
func (us DoctorSvc) DoctorDetailsEdit(ctx context.Context,r *doctorpb.DoctorDetailsEditRequest) (*doctorpb.DoctorDetailsEditResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditDoctorDetailsCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctorpb.DoctorDetailsEditResponse{
		ID:     int32(u.ID),
		Degree: u.Degree,
		Gender: u.Gender,
	},nil
}
// register doctor
func (us DoctorSvc) RegisterDoctorDetails(ctx context.Context,r *doctorpb.RegisterDoctorDetailsRequest) (*doctorpb.RegisterDoctorDetailsResponse, error) {
	fmt.Printf("%#v",r.DoctorTypeID)
	fmt.Printf("%#v",r.UserID)
	user := storage.DoctorDetails{
		UserID:       int(r.GetUserID()),
		DoctorTypeID: int(r.GetDoctorTypeID()),
		Degree:       r.GetDegree(),
		Gender:       r.GetGender(),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.RegisterDoctorDetailsCore(user)
	if err != nil {
		fmt.Println("response error", err.Error())
		return nil, err
	}

	return &doctorpb.RegisterDoctorDetailsResponse{
		User: &doctorpb.Doctor{
			UserID:       int32(u.UserID),
			DoctorTypeID: int32(u.DoctorTypeID),
			Degree:       u.Degree,
			Gender:       u.Gender,
		},
	}, nil
}

// register doctor schedule
func (us DoctorSvc) DoctorScheduleRegister(ctx context.Context, r *doctorpb.DoctorScheduleRegisterRequest) (*doctorpb.DoctorScheduleRegisterResponse, error) {
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
//list doctor
func (us DoctorSvc)DoctorList(ctx context.Context,r *doctorpb.DoctorListRequest) (*doctorpb.DoctorListResponse, error){
	user := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	u,err := us.core.ListDoctorCore(user)
	if err != nil {
		return nil,err
	}
	var totalusers []*doctorpb.User
	for _,value := range u {
		user:=&doctorpb.User{
			ID:        int32(value.ID),
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Username:  value.Username,
			Email:     value.Email,
			IsActive:  value.Is_active,
			Role:      value.Role,
		}
		totalusers = append(totalusers,user)
	}
	return &doctorpb.DoctorListResponse{
		User: totalusers,
	}, nil
}