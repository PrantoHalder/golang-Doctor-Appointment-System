package doctortype

import (
	"context"
	"fmt"

	doctortypepb "main.go/gunk/v1/doctortype"
	"main.go/usermgm/storage"
)

type CoreDoctorType interface {
	Registerdoctortype(storage.DoctorType) (*storage.DoctorType, error)
	EditDoctorType(u storage.Edit) (*storage.DoctorType, error)
	UpdateDoctorType(u storage.DoctorType) (*storage.DoctorType,error)
	DeleteDoctorTypeID(u storage.Edit) error
	ListDoctorType(u storage.UserFilter)([]storage.DoctorType,error)
}

type DoctorTypeSvc struct {
	doctortypepb.UnimplementedDoctorTypeServiceServer
	core CoreDoctorType
}

func NewDoctorTypeSvc(cu CoreDoctorType) *DoctorTypeSvc {
	return &DoctorTypeSvc{
		core: cu,
	}
}

//doctor type register
func (us DoctorTypeSvc) RegisterDoctorType(ctx context.Context, r *doctortypepb.RegisterDoctorTypeRequest) (*doctortypepb.RegisterDoctorTypeResponse, error) {
	user := storage.DoctorType{
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
	return &doctortypepb.RegisterDoctorTypeResponse{
		User: &doctortypepb.DoctorType{
			ID:         int32(u.ID),
			DoctorType: u.DoctorType,
		},
	}, nil
}
//edit doctor type
func (us DoctorTypeSvc) EditDoctorType(ctx context.Context,r *doctortypepb.EditDoctorTypeRequest) (*doctortypepb.EditDoctorTypeResponse, error) {
	user := storage.Edit{
		ID:         int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.EditDoctorType(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctortypepb.EditDoctorTypeResponse{
		User: &doctortypepb.DoctorType{
			ID: int32(u.ID),
			 DoctorType: u.DoctorType},
	}, nil

}
//update doctor type
func (us DoctorTypeSvc)UpdateDoctorType(ctx context.Context,r *doctortypepb.UpdateDoctorTypeRequest) (*doctortypepb.UpdateDoctorTypeResponse, error){
	user := storage.DoctorType{
		ID:         int(r.GetID()),
		DoctorType: r.GetDcotorType(),
		
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	u, err := us.core.UpdateDoctorType(user)
	if err != nil {
		fmt.Println("the error is in the serveice layer in Register after Register(user)")
		return nil, err
	}
	return &doctortypepb.UpdateDoctorTypeResponse{
		User: &doctortypepb.DoctorType{
			ID:         int32(u.ID),
			DoctorType: u.DoctorType,
		},
	}, nil
}
//delete doctor type
func (us DoctorTypeSvc)DeleteDoctorType(ctx context.Context,r *doctortypepb.DeleteDoctorTypeRequest) (*doctortypepb.DeleteDoctorTypeResponse, error) {
	user := storage.Edit{
		ID:         int(r.GetId()),
	}
	if err := user.Validate(); err != nil {
		fmt.Println("the error is in the serveice layer in Register after user.Validate")
		return nil, err
	}
	if err := us.core.DeleteDoctorTypeID(user);err != nil {
		return nil,err
	}
	return &doctortypepb.DeleteDoctorTypeResponse{},nil

}
//list doctor type
func (us DoctorTypeSvc)DoctorTypeList(ctx context.Context,r *doctortypepb.DoctorTypeListRequest) (*doctortypepb.DoctorTypeListResponse, error){
  user := storage.UserFilter{
  	SearchTerm: r.GetSearchTerm(),
  }
  u,err := us.core.ListDoctorType(user)
	if err != nil {
		return nil,err
	}
	var totaldoctortype []*doctortypepb.DoctorType
	for _,value := range u {
		user:=&doctortypepb.DoctorType{
			ID:         int32(value.ID),
			DoctorType: value.DoctorType,
		}
		totaldoctortype = append(totaldoctortype,user)
	}
   return &doctortypepb.DoctorTypeListResponse{
   	DoctorType:totaldoctortype ,
   },nil
}