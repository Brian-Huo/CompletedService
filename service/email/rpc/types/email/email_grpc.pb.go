// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: email.proto

package email

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

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailClient interface {
	Announcement(ctx context.Context, in *AnnouncementRequest, opts ...grpc.CallOption) (*AnnouncementResponse, error)
	InvoiceEmail(ctx context.Context, in *InvoiceEmailRequest, opts ...grpc.CallOption) (*InvoiceEmailResponse, error)
	GeneralEmail(ctx context.Context, in *GeneralEmailRequest, opts ...grpc.CallOption) (*GeneralEmailResponse, error)
}

type emailClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClient(cc grpc.ClientConnInterface) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) Announcement(ctx context.Context, in *AnnouncementRequest, opts ...grpc.CallOption) (*AnnouncementResponse, error) {
	out := new(AnnouncementResponse)
	err := c.cc.Invoke(ctx, "/email.Email/Announcement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) InvoiceEmail(ctx context.Context, in *InvoiceEmailRequest, opts ...grpc.CallOption) (*InvoiceEmailResponse, error) {
	out := new(InvoiceEmailResponse)
	err := c.cc.Invoke(ctx, "/email.Email/InvoiceEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) GeneralEmail(ctx context.Context, in *GeneralEmailRequest, opts ...grpc.CallOption) (*GeneralEmailResponse, error) {
	out := new(GeneralEmailResponse)
	err := c.cc.Invoke(ctx, "/email.Email/GeneralEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
// All implementations must embed UnimplementedEmailServer
// for forward compatibility
type EmailServer interface {
	Announcement(context.Context, *AnnouncementRequest) (*AnnouncementResponse, error)
	InvoiceEmail(context.Context, *InvoiceEmailRequest) (*InvoiceEmailResponse, error)
	GeneralEmail(context.Context, *GeneralEmailRequest) (*GeneralEmailResponse, error)
	mustEmbedUnimplementedEmailServer()
}

// UnimplementedEmailServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (UnimplementedEmailServer) Announcement(context.Context, *AnnouncementRequest) (*AnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Announcement not implemented")
}
func (UnimplementedEmailServer) InvoiceEmail(context.Context, *InvoiceEmailRequest) (*InvoiceEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvoiceEmail not implemented")
}
func (UnimplementedEmailServer) GeneralEmail(context.Context, *GeneralEmailRequest) (*GeneralEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneralEmail not implemented")
}
func (UnimplementedEmailServer) mustEmbedUnimplementedEmailServer() {}

// UnsafeEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServer will
// result in compilation errors.
type UnsafeEmailServer interface {
	mustEmbedUnimplementedEmailServer()
}

func RegisterEmailServer(s grpc.ServiceRegistrar, srv EmailServer) {
	s.RegisterService(&Email_ServiceDesc, srv)
}

func _Email_Announcement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Announcement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/Announcement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Announcement(ctx, req.(*AnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_InvoiceEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).InvoiceEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/InvoiceEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).InvoiceEmail(ctx, req.(*InvoiceEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_GeneralEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).GeneralEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/GeneralEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).GeneralEmail(ctx, req.(*GeneralEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Email_ServiceDesc is the grpc.ServiceDesc for Email service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Email_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Announcement",
			Handler:    _Email_Announcement_Handler,
		},
		{
			MethodName: "InvoiceEmail",
			Handler:    _Email_InvoiceEmail_Handler,
		},
		{
			MethodName: "GeneralEmail",
			Handler:    _Email_GeneralEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
