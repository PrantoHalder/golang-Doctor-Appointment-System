package user

import (
	"context"
	"fmt"

	userpb "main.go/gunk/v1/user"
	"main.go/usermgm/storage"
)

type CoreUser interface {
	Register(storage.Register) (*storage.Register, error)
	EditUserCore(storage.Edit) (*storage.User, error)
	UpdatePatient(storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteUserByID(storage.Edit) error
	ListUser(storage.UserFilter) ([]storage.User,error)
	RegisterAppointmentCore(storage.Appointment) (*storage.Appointment, error)
	EditStatusUserCore(us storage.Edit) (*storage.UpdateStatus, error)
	UpdateUserStatusCore(u storage.UpdateStatus) (*storage.UpdateStatus, error)
	ShowDoctorListToUserCore(us storage.Edit) ([]storage.ShowDoctorToPatient, error)
	AppinmentStatusCore(us storage.Edit) ([]storage.AppontmentStatus, error)
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
//appointment status
func (us UserSvc)AppoinmentStatus(ctx context.Context,r *userpb.AppoinmentStatusRequest) (*userpb.AppoinmentStatusResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}
	u, err := us.core.AppinmentStatusCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.EditUserCore(user.ID)")
		return nil, err
	}
    var totalusers []*userpb.AppontmentStatus
	for _,value := range u {
		user:=&userpb.AppontmentStatus{
			ID:           int32(value.ID),
			FirstName:    value.FirstName,
			LastName:     value.LastName,
			Is_Appointed: value.Is_Appointed,
			TimeSlot:     value.TimeSlot,
		}
		totalusers = append(totalusers,user)
	}


	return &userpb.AppoinmentStatusResponse{
		AppontmentStatus: totalusers,
	},nil
}
// show doctor list to patient
func (us UserSvc)ShowDoctorlistPatient(ctx context.Context,r *userpb.ShowDoctorlistPatientRequest) (*userpb.ShowDoctorlistPatientResponse, error) {
	user := storage.Edit{
		ID: int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}
	u, err := us.core.ShowDoctorListToUserCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.EditUserCore(user.ID)")
		return nil, err
	}
    var totalusers []*userpb.DoctorList
	for _,value := range u {
		user:=&userpb.DoctorList{
			ID:         int32(value.ID),
			FirstName:  value.FirstName,
			LastName:   value.LastName,
			Degree:     value.Degree,
			DoctorType: value.DoctorType,
			Gender:     value.Gender,
		}
		totalusers = append(totalusers,user)
	}


	return &userpb.ShowDoctorlistPatientResponse{
		DoctorList: totalusers,
	},nil
}
// user register
func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.Register{
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

// edit user
func (us UserSvc) UserEdit(cxt context.Context, r *userpb.UserEditRequest) (*userpb.UserEditResponse, error) {
	user := storage.Edit{
		ID: int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Login after login.Validate()")
		return nil, err
	}
	u, err := us.core.EditUserCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Login after us.core.EditUserCore(user.ID)")
		return nil, err
	}
	return &userpb.UserEditResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsActive:  u.Is_active,
		},
	}, nil
}
// edit user status
func (us UserSvc) EditPatientStatus(ctx context.Context,r *userpb.EditPatientStatusRequest) (*userpb.EditPatientStatusResponse, error)  {
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	u, err := us.core.EditStatusUserCore(user)
	if err != nil {
		return nil, err
	}
	return &userpb.EditPatientStatusResponse{
		ID:       int32(u.ID),
		IsActive: u.Is_active,
	},nil
		
}

// user update section
func (us UserSvc) UserUpdate(ctx context.Context, r *userpb.UserUpdateRequest) (*userpb.UserUpdateResponse, error) {
	user := storage.UpdateUser{
		ID:        int(r.GetID()),
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Is_active: r.GetIsActive(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.UpdatePatient(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.UserUpdateResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive:  u.Is_active,
		},
	}, nil
}
//update patient status
func (us UserSvc)UpdatePatientStatus(ctx context.Context,r *userpb.UpdatePatientStatusRequest) (*userpb.UpdatePatientStatusResponse, error){
	user := storage.UpdateUser{
		ID:        int(r.GetID()),
		Is_active: r.GetIsActive(),
	}
	u, err := us.core.UpdatePatient(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.UpdatePatientStatusResponse{
		IsActive:u.Is_active}, nil
}
//delete user
func (us UserSvc) UserDelete(ctx context.Context, r *userpb.UserDeleteRequest) (*userpb.UserDeleteResponse, error) {
	user := storage.Edit{
		ID: int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	if err := us.core.DeleteUserByID(user);err != nil{
		return nil,err
	}
    return &userpb.UserDeleteResponse{},nil
}
//list user
func (us UserSvc)UserList(ctx context.Context,r *userpb.UserlistRequest) (*userpb.UserListResponse, error){
	user := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	u,err := us.core.ListUser(user)
	if err != nil {
		return nil,err
	}
	var totalusers []*userpb.User
	for _,value := range u {
		user:=&userpb.User{
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
	return &userpb.UserListResponse{
		Users: totalusers,
	}, nil
  
}
// register appointment
func (us UserSvc) RegisterAppointment(ctx context.Context,r *userpb.RegisterAppointmentRequest) (*userpb.RegisterAppointmentResponse, error){
	user := storage.Appointment{
		UserID:          int(r.GetUserID()),
		DoctorDetailsID: int(r.GetDoctorDetailsID()),
		ScheduleID:      int(r.GetScheduleID()),
	}
	
	u, err := us.core.RegisterAppointmentCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &userpb.RegisterAppointmentResponse{
		Appointment: &userpb.Appointment{
			ID:              int32(u.ID),
			UserID:          int32(u.UserID),
			DoctorDetailsID: int32(u.DoctorDetailsID),
			ScheduleID:      int32(u.ScheduleID),
			Is_Appointed:    u.Is_Appointed,
		},
	}, nil
}