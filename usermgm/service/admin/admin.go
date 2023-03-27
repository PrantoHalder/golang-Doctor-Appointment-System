package admin

import (
	"context"
	"fmt"

	adminpb "main.go/gunk/v1/admin"
	"main.go/usermgm/storage"
)

type CoreAdmin interface {
	RegisterAdmin(storage.User) (*storage.User, error)
	RegisterDoctorAdminCore(storage.User) (*storage.User, error)
	RegisterPatient(storage.User) (*storage.User, error)
	EditAdminCore(storage.Edit) (*storage.User, error)
	UpdateAdminCore(storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteAdminByIDCore(u storage.Edit) error
	EditDoctorCore(storage.Edit) (*storage.User, error)
	UpdateDoctorCore(u storage.UpdateUser) (*storage.UpdateUser, error)
	DeleteDoctorByIDCore(u storage.Edit) error
	EditPatientCore(us storage.Edit) (*storage.User, error)
	UpdatePatientCore(u storage.UpdateUser) (*storage.UpdateUser, error)
	DeletePatientByIDCore(u storage.Edit) error
	ListAdminCore(u storage.UserFilter) ([]storage.User,error)
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
//doctor register by admin
func (us AdminSvc) RegisterDoctorAdmin(ctx context.Context,r *adminpb.RegisterDoctorAdminRequest) (*adminpb.RegisterDoctorAdminResponse, error) {
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
	u, err := us.core.RegisterDoctorAdminCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.RegisterDoctorAdminResponse{
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
// user register by admin
func (us AdminSvc) RegisterPatient(ctx context.Context, r *adminpb.RegisterPatientRequest) (*adminpb.RegisterPatientResponse, error) {
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
	u, err := us.core.RegisterPatient(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.RegisterPatientResponse{
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
// admin edit
func (us AdminSvc) AdminEdit(ctx context.Context,r *adminpb.AdminEditRequest) (*adminpb.AdminEditResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditAdminCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.AdminEditResponse{
		User: &adminpb.Edit{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//update admin
func (us AdminSvc) AdminUpdate(ctx context.Context,r *adminpb.AdminUpdateRequest) (*adminpb.AdminUpdateResponse, error) {
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
	u, err := us.core.UpdateAdminCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.AdminUpdateResponse{
		User: &adminpb.Update{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//delete admin
func (us AdminSvc) AdminDelete(ctx context.Context,r *adminpb.AdminDeleteRequest) (*adminpb.AdminDeleteResponse, error){
	user := storage.Edit{
		ID:        int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	err := us.core.DeleteAdminByIDCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.AdminDeleteResponse{},nil
		
}
//edit doctor
func (us AdminSvc)EditDoctorAdmin(ctx context.Context,r *adminpb.EditDoctorAdminRequest) (*adminpb.EditDoctorAdminResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditDoctorCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.EditDoctorAdminResponse{
		User: &adminpb.Edit{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//update doctor
func (us AdminSvc)UpdateDoctorAdmin(ctx context.Context,r *adminpb.UpdateDoctorAdminRequest) (*adminpb.UpdateDoctorAdminResponse, error){
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
	u, err := us.core.UpdateDoctorCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.UpdateDoctorAdminResponse{
		User: &adminpb.Update{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//delete doctor
func (us AdminSvc)DeleteDoctorByID(ctx context.Context,r *adminpb.DeleteAdminByIDRequest) (*adminpb.DeleteAdminByIDResponse, error){
	user := storage.Edit{
		ID:        int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	err := us.core.DeleteDoctorByIDCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.DeleteAdminByIDResponse{},nil
}
//edit patient
func (us AdminSvc)EditPatient(ctx context.Context,r *adminpb.EditPatientRequest) (*adminpb.EditPatientResponse, error){
	user := storage.Edit{
		ID: int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditPatientCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.EditPatientResponse{
		User: &adminpb.Edit{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//update patient
func (us AdminSvc)UpdatePatient(ctx context.Context,r *adminpb.UpdatePatientRequest) (*adminpb.UpdatePatientResponse, error){
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
	u, err := us.core.UpdatePatientCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.UpdatePatientResponse{
		User: &adminpb.Update{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			IsActive: u.Is_active,
		},
	}, nil
}
//delete a patient
func (us AdminSvc) DeletePatient(ctx context.Context,r *adminpb.DeletePatientRequest) (*adminpb.DeletePatientResponse, error){
	user := storage.Edit{
		ID:        int(r.GetID()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	err := us.core.DeletePatientByIDCore(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &adminpb.DeletePatientResponse{},nil
}
//list admin
func (us AdminSvc) AdminList(ctx context.Context,r *adminpb.AdminListRequest) (*adminpb.AdminListResponse, error){
	user := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	u,err := us.core.ListAdminCore(user)
	if err != nil {
		return nil,err
	}
	var totalusers []*adminpb.User
	for _,value := range u {
		user:=&adminpb.User{
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
	return &adminpb.AdminListResponse{
		Users: totalusers,
	}, nil
  
}