package doctortype

import (
	"context"
	"fmt"

	doctortypepb "main.go/gunk/v1/doctortype"
	"main.go/usermgm/storage"
)

type CoreDoctorType interface {
	Registerdoctortype(storage.Doctor_type) (*storage.Doctor_type, error)
}

type DoctorTypeSvc struct {
	doctortypepb.UnimplementedDoctorServiceServer
	core CoreDoctorType
}

func NewDoctorTypeSvc(cu CoreDoctorType) *DoctorTypeSvc {
	return &DoctorTypeSvc{
		core: cu,
	}
}

//doctor type register
func (us DoctorTypeSvc) RegisterDoctorType(ctx context.Context, r *doctortypepb.RegisterDoctorTypeRequest) (*doctortypepb.RegisterDoctorTypeResponse, error) {
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
	return &doctortypepb.RegisterDoctorTypeResponse{
		User: &doctortypepb.DoctorType{
			ID:         int32(u.ID),
			DoctorType: u.DoctorType,
		},
	}, nil
}