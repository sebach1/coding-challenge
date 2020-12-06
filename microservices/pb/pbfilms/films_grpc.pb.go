// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbfilms

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FilmsClient is the client API for Films service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmsClient interface {
	RetrieveFilm(ctx context.Context, in *RetrieveFilmRequest, opts ...grpc.CallOption) (*RetrieveFilmResponse, error)
	RetrieveFilms(ctx context.Context, in *RetrieveFilmsRequest, opts ...grpc.CallOption) (*RetrieveFilmsResponse, error)
	RetrieveFilmsWithPeople(ctx context.Context, in *RetrieveFilmsWithPeopleRequest, opts ...grpc.CallOption) (*RetrieveFilmsWithPeopleResponse, error)
	CreateFilms(ctx context.Context, opts ...grpc.CallOption) (Films_CreateFilmsClient, error)
	CreatePeople(ctx context.Context, opts ...grpc.CallOption) (Films_CreatePeopleClient, error)
	CreateJoinPeopleFilm(ctx context.Context, opts ...grpc.CallOption) (Films_CreateJoinPeopleFilmClient, error)
}

type filmsClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmsClient(cc grpc.ClientConnInterface) FilmsClient {
	return &filmsClient{cc}
}

func (c *filmsClient) RetrieveFilm(ctx context.Context, in *RetrieveFilmRequest, opts ...grpc.CallOption) (*RetrieveFilmResponse, error) {
	out := new(RetrieveFilmResponse)
	err := c.cc.Invoke(ctx, "/films.Films/RetrieveFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) RetrieveFilms(ctx context.Context, in *RetrieveFilmsRequest, opts ...grpc.CallOption) (*RetrieveFilmsResponse, error) {
	out := new(RetrieveFilmsResponse)
	err := c.cc.Invoke(ctx, "/films.Films/RetrieveFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) RetrieveFilmsWithPeople(ctx context.Context, in *RetrieveFilmsWithPeopleRequest, opts ...grpc.CallOption) (*RetrieveFilmsWithPeopleResponse, error) {
	out := new(RetrieveFilmsWithPeopleResponse)
	err := c.cc.Invoke(ctx, "/films.Films/RetrieveFilmsWithPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) CreateFilms(ctx context.Context, opts ...grpc.CallOption) (Films_CreateFilmsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Films_serviceDesc.Streams[0], "/films.Films/CreateFilms", opts...)
	if err != nil {
		return nil, err
	}
	x := &filmsCreateFilmsClient{stream}
	return x, nil
}

type Films_CreateFilmsClient interface {
	Send(*CreateFilmRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type filmsCreateFilmsClient struct {
	grpc.ClientStream
}

func (x *filmsCreateFilmsClient) Send(m *CreateFilmRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filmsCreateFilmsClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filmsClient) CreatePeople(ctx context.Context, opts ...grpc.CallOption) (Films_CreatePeopleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Films_serviceDesc.Streams[1], "/films.Films/CreatePeople", opts...)
	if err != nil {
		return nil, err
	}
	x := &filmsCreatePeopleClient{stream}
	return x, nil
}

type Films_CreatePeopleClient interface {
	Send(*CreatePeopleRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type filmsCreatePeopleClient struct {
	grpc.ClientStream
}

func (x *filmsCreatePeopleClient) Send(m *CreatePeopleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filmsCreatePeopleClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filmsClient) CreateJoinPeopleFilm(ctx context.Context, opts ...grpc.CallOption) (Films_CreateJoinPeopleFilmClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Films_serviceDesc.Streams[2], "/films.Films/CreateJoinPeopleFilm", opts...)
	if err != nil {
		return nil, err
	}
	x := &filmsCreateJoinPeopleFilmClient{stream}
	return x, nil
}

type Films_CreateJoinPeopleFilmClient interface {
	Send(*CreateJoinPeopleFilmRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type filmsCreateJoinPeopleFilmClient struct {
	grpc.ClientStream
}

func (x *filmsCreateJoinPeopleFilmClient) Send(m *CreateJoinPeopleFilmRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filmsCreateJoinPeopleFilmClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilmsServer is the server API for Films service.
// All implementations must embed UnimplementedFilmsServer
// for forward compatibility
type FilmsServer interface {
	RetrieveFilm(context.Context, *RetrieveFilmRequest) (*RetrieveFilmResponse, error)
	RetrieveFilms(context.Context, *RetrieveFilmsRequest) (*RetrieveFilmsResponse, error)
	RetrieveFilmsWithPeople(context.Context, *RetrieveFilmsWithPeopleRequest) (*RetrieveFilmsWithPeopleResponse, error)
	CreateFilms(Films_CreateFilmsServer) error
	CreatePeople(Films_CreatePeopleServer) error
	CreateJoinPeopleFilm(Films_CreateJoinPeopleFilmServer) error
	mustEmbedUnimplementedFilmsServer()
}

// UnimplementedFilmsServer must be embedded to have forward compatible implementations.
type UnimplementedFilmsServer struct {
}

func (UnimplementedFilmsServer) RetrieveFilm(context.Context, *RetrieveFilmRequest) (*RetrieveFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveFilm not implemented")
}
func (UnimplementedFilmsServer) RetrieveFilms(context.Context, *RetrieveFilmsRequest) (*RetrieveFilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveFilms not implemented")
}
func (UnimplementedFilmsServer) RetrieveFilmsWithPeople(context.Context, *RetrieveFilmsWithPeopleRequest) (*RetrieveFilmsWithPeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveFilmsWithPeople not implemented")
}
func (UnimplementedFilmsServer) CreateFilms(Films_CreateFilmsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateFilms not implemented")
}
func (UnimplementedFilmsServer) CreatePeople(Films_CreatePeopleServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePeople not implemented")
}
func (UnimplementedFilmsServer) CreateJoinPeopleFilm(Films_CreateJoinPeopleFilmServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateJoinPeopleFilm not implemented")
}
func (UnimplementedFilmsServer) mustEmbedUnimplementedFilmsServer() {}

// UnsafeFilmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmsServer will
// result in compilation errors.
type UnsafeFilmsServer interface {
	mustEmbedUnimplementedFilmsServer()
}

func RegisterFilmsServer(s grpc.ServiceRegistrar, srv FilmsServer) {
	s.RegisterService(&_Films_serviceDesc, srv)
}

func _Films_RetrieveFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).RetrieveFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.Films/RetrieveFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).RetrieveFilm(ctx, req.(*RetrieveFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_RetrieveFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).RetrieveFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.Films/RetrieveFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).RetrieveFilms(ctx, req.(*RetrieveFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_RetrieveFilmsWithPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveFilmsWithPeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).RetrieveFilmsWithPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.Films/RetrieveFilmsWithPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).RetrieveFilmsWithPeople(ctx, req.(*RetrieveFilmsWithPeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_CreateFilms_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilmsServer).CreateFilms(&filmsCreateFilmsServer{stream})
}

type Films_CreateFilmsServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*CreateFilmRequest, error)
	grpc.ServerStream
}

type filmsCreateFilmsServer struct {
	grpc.ServerStream
}

func (x *filmsCreateFilmsServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filmsCreateFilmsServer) Recv() (*CreateFilmRequest, error) {
	m := new(CreateFilmRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Films_CreatePeople_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilmsServer).CreatePeople(&filmsCreatePeopleServer{stream})
}

type Films_CreatePeopleServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*CreatePeopleRequest, error)
	grpc.ServerStream
}

type filmsCreatePeopleServer struct {
	grpc.ServerStream
}

func (x *filmsCreatePeopleServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filmsCreatePeopleServer) Recv() (*CreatePeopleRequest, error) {
	m := new(CreatePeopleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Films_CreateJoinPeopleFilm_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilmsServer).CreateJoinPeopleFilm(&filmsCreateJoinPeopleFilmServer{stream})
}

type Films_CreateJoinPeopleFilmServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*CreateJoinPeopleFilmRequest, error)
	grpc.ServerStream
}

type filmsCreateJoinPeopleFilmServer struct {
	grpc.ServerStream
}

func (x *filmsCreateJoinPeopleFilmServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filmsCreateJoinPeopleFilmServer) Recv() (*CreateJoinPeopleFilmRequest, error) {
	m := new(CreateJoinPeopleFilmRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Films_serviceDesc = grpc.ServiceDesc{
	ServiceName: "films.Films",
	HandlerType: (*FilmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveFilm",
			Handler:    _Films_RetrieveFilm_Handler,
		},
		{
			MethodName: "RetrieveFilms",
			Handler:    _Films_RetrieveFilms_Handler,
		},
		{
			MethodName: "RetrieveFilmsWithPeople",
			Handler:    _Films_RetrieveFilmsWithPeople_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateFilms",
			Handler:       _Films_CreateFilms_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreatePeople",
			Handler:       _Films_CreatePeople_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateJoinPeopleFilm",
			Handler:       _Films_CreateJoinPeopleFilm_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "films.proto",
}
