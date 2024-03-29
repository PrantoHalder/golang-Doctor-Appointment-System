// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doctortypepb

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

// DoctorTypeServiceClient is the client API for DoctorTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorTypeServiceClient interface {
	RegisterDoctorType(ctx context.Context, in *RegisterDoctorTypeRequest, opts ...grpc.CallOption) (*RegisterDoctorTypeResponse, error)
	EditDoctorType(ctx context.Context, in *EditDoctorTypeRequest, opts ...grpc.CallOption) (*EditDoctorTypeResponse, error)
	UpdateDoctorType(ctx context.Context, in *UpdateDoctorTypeRequest, opts ...grpc.CallOption) (*UpdateDoctorTypeResponse, error)
	DeleteDoctorType(ctx context.Context, in *DeleteDoctorTypeRequest, opts ...grpc.CallOption) (*DeleteDoctorTypeResponse, error)
	DoctorTypeList(ctx context.Context, in *DoctorTypeListRequest, opts ...grpc.CallOption) (*DoctorTypeListResponse, error)
}

type doctorTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorTypeServiceClient(cc grpc.ClientConnInterface) DoctorTypeServiceClient {
	return &doctorTypeServiceClient{cc}
}

func (c *doctorTypeServiceClient) RegisterDoctorType(ctx context.Context, in *RegisterDoctorTypeRequest, opts ...grpc.CallOption) (*RegisterDoctorTypeResponse, error) {
	out := new(RegisterDoctorTypeResponse)
	err := c.cc.Invoke(ctx, "/doctortypepb.DoctorTypeService/RegisterDoctorType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorTypeServiceClient) EditDoctorType(ctx context.Context, in *EditDoctorTypeRequest, opts ...grpc.CallOption) (*EditDoctorTypeResponse, error) {
	out := new(EditDoctorTypeResponse)
	err := c.cc.Invoke(ctx, "/doctortypepb.DoctorTypeService/EditDoctorType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorTypeServiceClient) UpdateDoctorType(ctx context.Context, in *UpdateDoctorTypeRequest, opts ...grpc.CallOption) (*UpdateDoctorTypeResponse, error) {
	out := new(UpdateDoctorTypeResponse)
	err := c.cc.Invoke(ctx, "/doctortypepb.DoctorTypeService/UpdateDoctorType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorTypeServiceClient) DeleteDoctorType(ctx context.Context, in *DeleteDoctorTypeRequest, opts ...grpc.CallOption) (*DeleteDoctorTypeResponse, error) {
	out := new(DeleteDoctorTypeResponse)
	err := c.cc.Invoke(ctx, "/doctortypepb.DoctorTypeService/DeleteDoctorType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorTypeServiceClient) DoctorTypeList(ctx context.Context, in *DoctorTypeListRequest, opts ...grpc.CallOption) (*DoctorTypeListResponse, error) {
	out := new(DoctorTypeListResponse)
	err := c.cc.Invoke(ctx, "/doctortypepb.DoctorTypeService/DoctorTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorTypeServiceServer is the server API for DoctorTypeService service.
// All implementations must embed UnimplementedDoctorTypeServiceServer
// for forward compatibility
type DoctorTypeServiceServer interface {
	RegisterDoctorType(context.Context, *RegisterDoctorTypeRequest) (*RegisterDoctorTypeResponse, error)
	EditDoctorType(context.Context, *EditDoctorTypeRequest) (*EditDoctorTypeResponse, error)
	UpdateDoctorType(context.Context, *UpdateDoctorTypeRequest) (*UpdateDoctorTypeResponse, error)
	DeleteDoctorType(context.Context, *DeleteDoctorTypeRequest) (*DeleteDoctorTypeResponse, error)
	DoctorTypeList(context.Context, *DoctorTypeListRequest) (*DoctorTypeListResponse, error)
	mustEmbedUnimplementedDoctorTypeServiceServer()
}

// UnimplementedDoctorTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorTypeServiceServer struct {
}

func (UnimplementedDoctorTypeServiceServer) RegisterDoctorType(context.Context, *RegisterDoctorTypeRequest) (*RegisterDoctorTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctorType not implemented")
}
func (UnimplementedDoctorTypeServiceServer) EditDoctorType(context.Context, *EditDoctorTypeRequest) (*EditDoctorTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDoctorType not implemented")
}
func (UnimplementedDoctorTypeServiceServer) UpdateDoctorType(context.Context, *UpdateDoctorTypeRequest) (*UpdateDoctorTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctorType not implemented")
}
func (UnimplementedDoctorTypeServiceServer) DeleteDoctorType(context.Context, *DeleteDoctorTypeRequest) (*DeleteDoctorTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctorType not implemented")
}
func (UnimplementedDoctorTypeServiceServer) DoctorTypeList(context.Context, *DoctorTypeListRequest) (*DoctorTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorTypeList not implemented")
}
func (UnimplementedDoctorTypeServiceServer) mustEmbedUnimplementedDoctorTypeServiceServer() {}

// UnsafeDoctorTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorTypeServiceServer will
// result in compilation errors.
type UnsafeDoctorTypeServiceServer interface {
	mustEmbedUnimplementedDoctorTypeServiceServer()
}

func RegisterDoctorTypeServiceServer(s grpc.ServiceRegistrar, srv DoctorTypeServiceServer) {
	s.RegisterService(&DoctorTypeService_ServiceDesc, srv)
}

func _DoctorTypeService_RegisterDoctorType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDoctorTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorTypeServiceServer).RegisterDoctorType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctortypepb.DoctorTypeService/RegisterDoctorType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorTypeServiceServer).RegisterDoctorType(ctx, req.(*RegisterDoctorTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorTypeService_EditDoctorType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditDoctorTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorTypeServiceServer).EditDoctorType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctortypepb.DoctorTypeService/EditDoctorType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorTypeServiceServer).EditDoctorType(ctx, req.(*EditDoctorTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorTypeService_UpdateDoctorType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDoctorTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorTypeServiceServer).UpdateDoctorType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctortypepb.DoctorTypeService/UpdateDoctorType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorTypeServiceServer).UpdateDoctorType(ctx, req.(*UpdateDoctorTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorTypeService_DeleteDoctorType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDoctorTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorTypeServiceServer).DeleteDoctorType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctortypepb.DoctorTypeService/DeleteDoctorType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorTypeServiceServer).DeleteDoctorType(ctx, req.(*DeleteDoctorTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorTypeService_DoctorTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorTypeServiceServer).DoctorTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctortypepb.DoctorTypeService/DoctorTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorTypeServiceServer).DoctorTypeList(ctx, req.(*DoctorTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorTypeService_ServiceDesc is the grpc.ServiceDesc for DoctorTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctortypepb.DoctorTypeService",
	HandlerType: (*DoctorTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDoctorType",
			Handler:    _DoctorTypeService_RegisterDoctorType_Handler,
		},
		{
			MethodName: "EditDoctorType",
			Handler:    _DoctorTypeService_EditDoctorType_Handler,
		},
		{
			MethodName: "UpdateDoctorType",
			Handler:    _DoctorTypeService_UpdateDoctorType_Handler,
		},
		{
			MethodName: "DeleteDoctorType",
			Handler:    _DoctorTypeService_DeleteDoctorType_Handler,
		},
		{
			MethodName: "DoctorTypeList",
			Handler:    _DoctorTypeService_DoctorTypeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.go/gunk/v1/doctortype/all.proto",
}
