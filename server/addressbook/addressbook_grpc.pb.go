// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: addressbook.proto

package addressbook

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AddressBookService_ListPerson_FullMethodName = "/addressbook.AddressBookService/ListPerson"
	AddressBookService_AddPerson_FullMethodName  = "/addressbook.AddressBookService/AddPerson"
)

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// service is the keyword for rpc which is for implementing api design
type AddressBookServiceClient interface {
	ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (*ListPersonResponse, error)
	AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*AddPersonResponse, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (*ListPersonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPersonResponse)
	err := c.cc.Invoke(ctx, AddressBookService_ListPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*AddPersonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPersonResponse)
	err := c.cc.Invoke(ctx, AddressBookService_AddPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility.
//
// service is the keyword for rpc which is for implementing api design
type AddressBookServiceServer interface {
	ListPerson(context.Context, *ListPersonRequest) (*ListPersonResponse, error)
	AddPerson(context.Context, *Person) (*AddPersonResponse, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAddressBookServiceServer struct{}

func (UnimplementedAddressBookServiceServer) ListPerson(context.Context, *ListPersonRequest) (*ListPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerson not implemented")
}
func (UnimplementedAddressBookServiceServer) AddPerson(context.Context, *Person) (*AddPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}
func (UnimplementedAddressBookServiceServer) testEmbeddedByValue()                            {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	// If the following call pancis, it indicates UnimplementedAddressBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_ListPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).ListPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressBookService_ListPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).ListPerson(ctx, req.(*ListPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressBookService_AddPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).AddPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "addressbook.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPerson",
			Handler:    _AddressBookService_ListPerson_Handler,
		},
		{
			MethodName: "AddPerson",
			Handler:    _AddressBookService_AddPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addressbook.proto",
}
