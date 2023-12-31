// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: protocols/books.proto

package protocols

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

const (
	BookServices_GetByAuthor_FullMethodName = "/protocols.BookServices/getByAuthor"
	BookServices_GetByName_FullMethodName   = "/protocols.BookServices/getByName"
)

// BookServicesClient is the client API for BookServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServicesClient interface {
	GetByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*NameResponse, error)
	GetByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Author, error)
}

type bookServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServicesClient(cc grpc.ClientConnInterface) BookServicesClient {
	return &bookServicesClient{cc}
}

func (c *bookServicesClient) GetByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, BookServices_GetByAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServicesClient) GetByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, BookServices_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServicesServer is the server API for BookServices service.
// All implementations must embed UnimplementedBookServicesServer
// for forward compatibility
type BookServicesServer interface {
	GetByAuthor(context.Context, *Author) (*NameResponse, error)
	GetByName(context.Context, *Name) (*Author, error)
	mustEmbedUnimplementedBookServicesServer()
}

// UnimplementedBookServicesServer must be embedded to have forward compatible implementations.
type UnimplementedBookServicesServer struct {
}

func (UnimplementedBookServicesServer) GetByAuthor(context.Context, *Author) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAuthor not implemented")
}
func (UnimplementedBookServicesServer) GetByName(context.Context, *Name) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedBookServicesServer) mustEmbedUnimplementedBookServicesServer() {}

// UnsafeBookServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServicesServer will
// result in compilation errors.
type UnsafeBookServicesServer interface {
	mustEmbedUnimplementedBookServicesServer()
}

func RegisterBookServicesServer(s grpc.ServiceRegistrar, srv BookServicesServer) {
	s.RegisterService(&BookServices_ServiceDesc, srv)
}

func _BookServices_GetByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServicesServer).GetByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookServices_GetByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServicesServer).GetByAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookServices_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServicesServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookServices_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServicesServer).GetByName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

// BookServices_ServiceDesc is the grpc.ServiceDesc for BookServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.BookServices",
	HandlerType: (*BookServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByAuthor",
			Handler:    _BookServices_GetByAuthor_Handler,
		},
		{
			MethodName: "getByName",
			Handler:    _BookServices_GetByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocols/books.proto",
}
