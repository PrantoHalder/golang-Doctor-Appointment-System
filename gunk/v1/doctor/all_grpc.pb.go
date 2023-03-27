// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doctorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	RegisterDoctorDetails(ctx context.Context, in *RegisterDoctorDetailsRequest, opts ...grpc.CallOption) (*RegisterDoctorDetailsResponse, error)
	DoctorScheduleRegister(ctx context.Context, in *DoctorScheduleRegisterRequest, opts ...grpc.CallOption) (*DoctorScheduleRegisterResponse, error)
	DoctorScheduleEdit(ctx context.Context, in *DoctorScheduleEditRequest, opts ...grpc.CallOption) (*DoctorScheduleEditResponse, error)
	DoctorScheduleUpdate(ctx context.Context, in *DoctorScheduleUpdateRequest, opts ...grpc.CallOption) (*DoctorScheduleUpdateResponse, error)
	DoctorList(ctx context.Context, in *DoctorListRequest, opts ...grpc.CallOption) (*DoctorListResponse, error)
	EditDoctorStatus(ctx context.Context, in *EditDoctorStatusRequest, opts ...grpc.CallOption) (*EditDoctorStatusResponse, error)
	UpdateDoctorStatus(ctx context.Context, in *UpdateDoctorStatusRequest, opts ...grpc.CallOption) (*UpdateDoctorStatusResponse, error)
	DoctorDetailsEdit(ctx context.Context, in *DoctorDetailsEditRequest, opts ...grpc.CallOption) (*DoctorDetailsEditResponse, error)
	DoctorDetailsUpdate(ctx context.Context, in *DoctorDetailsUpdateRequest, opts ...grpc.CallOption) (*DoctorDetailsUpdateResponse, error)
	ApproveAppointmentEdit(ctx context.Context, in *ApproveAppointmentEditRequest, opts ...grpc.CallOption) (*ApproveAppointmentEditResponse, error)
	ApproveAppointmentUpdate(ctx context.Context, in *ApproveAppointmentUpdateRequest, opts ...grpc.CallOption) (*ApproveAppointmentUpdateResponse, error)
	DoctorDetailsList(ctx context.Context, in *DoctorDetailsListRequest, opts ...grpc.CallOption) (*DoctorDetailsListResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) RegisterDoctorDetails(ctx context.Context, in *RegisterDoctorDetailsRequest, opts ...grpc.CallOption) (*RegisterDoctorDetailsResponse, error) {
	out := new(RegisterDoctorDetailsResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/RegisterDoctorDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorScheduleRegister(ctx context.Context, in *DoctorScheduleRegisterRequest, opts ...grpc.CallOption) (*DoctorScheduleRegisterResponse, error) {
	out := new(DoctorScheduleRegisterResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorScheduleRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorScheduleEdit(ctx context.Context, in *DoctorScheduleEditRequest, opts ...grpc.CallOption) (*DoctorScheduleEditResponse, error) {
	out := new(DoctorScheduleEditResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorScheduleEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorScheduleUpdate(ctx context.Context, in *DoctorScheduleUpdateRequest, opts ...grpc.CallOption) (*DoctorScheduleUpdateResponse, error) {
	out := new(DoctorScheduleUpdateResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorScheduleUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorList(ctx context.Context, in *DoctorListRequest, opts ...grpc.CallOption) (*DoctorListResponse, error) {
	out := new(DoctorListResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) EditDoctorStatus(ctx context.Context, in *EditDoctorStatusRequest, opts ...grpc.CallOption) (*EditDoctorStatusResponse, error) {
	out := new(EditDoctorStatusResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/EditDoctorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateDoctorStatus(ctx context.Context, in *UpdateDoctorStatusRequest, opts ...grpc.CallOption) (*UpdateDoctorStatusResponse, error) {
	out := new(UpdateDoctorStatusResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/UpdateDoctorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorDetailsEdit(ctx context.Context, in *DoctorDetailsEditRequest, opts ...grpc.CallOption) (*DoctorDetailsEditResponse, error) {
	out := new(DoctorDetailsEditResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorDetailsEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorDetailsUpdate(ctx context.Context, in *DoctorDetailsUpdateRequest, opts ...grpc.CallOption) (*DoctorDetailsUpdateResponse, error) {
	out := new(DoctorDetailsUpdateResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorDetailsUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ApproveAppointmentEdit(ctx context.Context, in *ApproveAppointmentEditRequest, opts ...grpc.CallOption) (*ApproveAppointmentEditResponse, error) {
	out := new(ApproveAppointmentEditResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/ApproveAppointmentEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ApproveAppointmentUpdate(ctx context.Context, in *ApproveAppointmentUpdateRequest, opts ...grpc.CallOption) (*ApproveAppointmentUpdateResponse, error) {
	out := new(ApproveAppointmentUpdateResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/ApproveAppointmentUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorDetailsList(ctx context.Context, in *DoctorDetailsListRequest, opts ...grpc.CallOption) (*DoctorDetailsListResponse, error) {
	out := new(DoctorDetailsListResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorDetailsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	RegisterDoctorDetails(context.Context, *RegisterDoctorDetailsRequest) (*RegisterDoctorDetailsResponse, error)
	DoctorScheduleRegister(context.Context, *DoctorScheduleRegisterRequest) (*DoctorScheduleRegisterResponse, error)
	DoctorScheduleEdit(context.Context, *DoctorScheduleEditRequest) (*DoctorScheduleEditResponse, error)
	DoctorScheduleUpdate(context.Context, *DoctorScheduleUpdateRequest) (*DoctorScheduleUpdateResponse, error)
	DoctorList(context.Context, *DoctorListRequest) (*DoctorListResponse, error)
	EditDoctorStatus(context.Context, *EditDoctorStatusRequest) (*EditDoctorStatusResponse, error)
	UpdateDoctorStatus(context.Context, *UpdateDoctorStatusRequest) (*UpdateDoctorStatusResponse, error)
	DoctorDetailsEdit(context.Context, *DoctorDetailsEditRequest) (*DoctorDetailsEditResponse, error)
	DoctorDetailsUpdate(context.Context, *DoctorDetailsUpdateRequest) (*DoctorDetailsUpdateResponse, error)
	ApproveAppointmentEdit(context.Context, *ApproveAppointmentEditRequest) (*ApproveAppointmentEditResponse, error)
	ApproveAppointmentUpdate(context.Context, *ApproveAppointmentUpdateRequest) (*ApproveAppointmentUpdateResponse, error)
	DoctorDetailsList(context.Context, *DoctorDetailsListRequest) (*DoctorDetailsListResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) RegisterDoctorDetails(context.Context, *RegisterDoctorDetailsRequest) (*RegisterDoctorDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctorDetails not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorScheduleRegister(context.Context, *DoctorScheduleRegisterRequest) (*DoctorScheduleRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorScheduleRegister not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorScheduleEdit(context.Context, *DoctorScheduleEditRequest) (*DoctorScheduleEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorScheduleEdit not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorScheduleUpdate(context.Context, *DoctorScheduleUpdateRequest) (*DoctorScheduleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorScheduleUpdate not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorList(context.Context, *DoctorListRequest) (*DoctorListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorList not implemented")
}
func (UnimplementedDoctorServiceServer) EditDoctorStatus(context.Context, *EditDoctorStatusRequest) (*EditDoctorStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDoctorStatus not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateDoctorStatus(context.Context, *UpdateDoctorStatusRequest) (*UpdateDoctorStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctorStatus not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorDetailsEdit(context.Context, *DoctorDetailsEditRequest) (*DoctorDetailsEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorDetailsEdit not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorDetailsUpdate(context.Context, *DoctorDetailsUpdateRequest) (*DoctorDetailsUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorDetailsUpdate not implemented")
}
func (UnimplementedDoctorServiceServer) ApproveAppointmentEdit(context.Context, *ApproveAppointmentEditRequest) (*ApproveAppointmentEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAppointmentEdit not implemented")
}
func (UnimplementedDoctorServiceServer) ApproveAppointmentUpdate(context.Context, *ApproveAppointmentUpdateRequest) (*ApproveAppointmentUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAppointmentUpdate not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorDetailsList(context.Context, *DoctorDetailsListRequest) (*DoctorDetailsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorDetailsList not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_RegisterDoctorDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDoctorDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).RegisterDoctorDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/RegisterDoctorDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).RegisterDoctorDetails(ctx, req.(*RegisterDoctorDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorScheduleRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorScheduleRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorScheduleRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorScheduleRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorScheduleRegister(ctx, req.(*DoctorScheduleRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorScheduleEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorScheduleEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorScheduleEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorScheduleEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorScheduleEdit(ctx, req.(*DoctorScheduleEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorScheduleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorScheduleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorScheduleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorScheduleUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorScheduleUpdate(ctx, req.(*DoctorScheduleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorList(ctx, req.(*DoctorListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_EditDoctorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditDoctorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).EditDoctorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/EditDoctorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).EditDoctorStatus(ctx, req.(*EditDoctorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UpdateDoctorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDoctorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UpdateDoctorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/UpdateDoctorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UpdateDoctorStatus(ctx, req.(*UpdateDoctorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorDetailsEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorDetailsEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorDetailsEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorDetailsEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorDetailsEdit(ctx, req.(*DoctorDetailsEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorDetailsUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorDetailsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorDetailsUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorDetailsUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorDetailsUpdate(ctx, req.(*DoctorDetailsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ApproveAppointmentEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveAppointmentEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ApproveAppointmentEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/ApproveAppointmentEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ApproveAppointmentEdit(ctx, req.(*ApproveAppointmentEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ApproveAppointmentUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveAppointmentUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ApproveAppointmentUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/ApproveAppointmentUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ApproveAppointmentUpdate(ctx, req.(*ApproveAppointmentUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorDetailsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorDetailsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorDetailsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorDetailsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorDetailsList(ctx, req.(*DoctorDetailsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctorpb.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDoctorDetails",
			Handler:    _DoctorService_RegisterDoctorDetails_Handler,
		},
		{
			MethodName: "DoctorScheduleRegister",
			Handler:    _DoctorService_DoctorScheduleRegister_Handler,
		},
		{
			MethodName: "DoctorScheduleEdit",
			Handler:    _DoctorService_DoctorScheduleEdit_Handler,
		},
		{
			MethodName: "DoctorScheduleUpdate",
			Handler:    _DoctorService_DoctorScheduleUpdate_Handler,
		},
		{
			MethodName: "DoctorList",
			Handler:    _DoctorService_DoctorList_Handler,
		},
		{
			MethodName: "EditDoctorStatus",
			Handler:    _DoctorService_EditDoctorStatus_Handler,
		},
		{
			MethodName: "UpdateDoctorStatus",
			Handler:    _DoctorService_UpdateDoctorStatus_Handler,
		},
		{
			MethodName: "DoctorDetailsEdit",
			Handler:    _DoctorService_DoctorDetailsEdit_Handler,
		},
		{
			MethodName: "DoctorDetailsUpdate",
			Handler:    _DoctorService_DoctorDetailsUpdate_Handler,
		},
		{
			MethodName: "ApproveAppointmentEdit",
			Handler:    _DoctorService_ApproveAppointmentEdit_Handler,
		},
		{
			MethodName: "ApproveAppointmentUpdate",
			Handler:    _DoctorService_ApproveAppointmentUpdate_Handler,
		},
		{
			MethodName: "DoctorDetailsList",
			Handler:    _DoctorService_DoctorDetailsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.go/gunk/v1/doctor/all.proto",
}
