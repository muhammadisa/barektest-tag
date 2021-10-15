// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BareksaNewsServiceClient is the client API for BareksaNewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BareksaNewsServiceClient interface {
	AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTag(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTags(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Tags, error)
	AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTopic(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTopics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Topics, error)
	AddNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteNews(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetNewses(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Newses, error)
}

type bareksaNewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBareksaNewsServiceClient(cc grpc.ClientConnInterface) BareksaNewsServiceClient {
	return &bareksaNewsServiceClient{cc}
}

func (c *bareksaNewsServiceClient) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/EditTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) DeleteTag(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) GetTags(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/AddTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) EditTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/EditTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) DeleteTopic(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) GetTopics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Topics, error) {
	out := new(Topics)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) AddNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/AddNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) EditNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/EditNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) DeleteNews(ctx context.Context, in *Select, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/DeleteNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bareksaNewsServiceClient) GetNewses(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Newses, error) {
	out := new(Newses)
	err := c.cc.Invoke(ctx, "/api.v1.BareksaNewsService/GetNewses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BareksaNewsServiceServer is the server API for BareksaNewsService service.
// All implementations should embed UnimplementedBareksaNewsServiceServer
// for forward compatibility
type BareksaNewsServiceServer interface {
	AddTag(context.Context, *Tag) (*emptypb.Empty, error)
	EditTag(context.Context, *Tag) (*emptypb.Empty, error)
	DeleteTag(context.Context, *Select) (*emptypb.Empty, error)
	GetTags(context.Context, *emptypb.Empty) (*Tags, error)
	AddTopic(context.Context, *Topic) (*emptypb.Empty, error)
	EditTopic(context.Context, *Topic) (*emptypb.Empty, error)
	DeleteTopic(context.Context, *Select) (*emptypb.Empty, error)
	GetTopics(context.Context, *emptypb.Empty) (*Topics, error)
	AddNews(context.Context, *News) (*emptypb.Empty, error)
	EditNews(context.Context, *News) (*emptypb.Empty, error)
	DeleteNews(context.Context, *Select) (*emptypb.Empty, error)
	GetNewses(context.Context, *Filters) (*Newses, error)
}

// UnimplementedBareksaNewsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBareksaNewsServiceServer struct {
}

func (UnimplementedBareksaNewsServiceServer) AddTag(context.Context, *Tag) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedBareksaNewsServiceServer) EditTag(context.Context, *Tag) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTag not implemented")
}
func (UnimplementedBareksaNewsServiceServer) DeleteTag(context.Context, *Select) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedBareksaNewsServiceServer) GetTags(context.Context, *emptypb.Empty) (*Tags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedBareksaNewsServiceServer) AddTopic(context.Context, *Topic) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopic not implemented")
}
func (UnimplementedBareksaNewsServiceServer) EditTopic(context.Context, *Topic) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTopic not implemented")
}
func (UnimplementedBareksaNewsServiceServer) DeleteTopic(context.Context, *Select) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedBareksaNewsServiceServer) GetTopics(context.Context, *emptypb.Empty) (*Topics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (UnimplementedBareksaNewsServiceServer) AddNews(context.Context, *News) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNews not implemented")
}
func (UnimplementedBareksaNewsServiceServer) EditNews(context.Context, *News) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditNews not implemented")
}
func (UnimplementedBareksaNewsServiceServer) DeleteNews(context.Context, *Select) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNews not implemented")
}
func (UnimplementedBareksaNewsServiceServer) GetNewses(context.Context, *Filters) (*Newses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewses not implemented")
}

// UnsafeBareksaNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BareksaNewsServiceServer will
// result in compilation errors.
type UnsafeBareksaNewsServiceServer interface {
	mustEmbedUnimplementedBareksaNewsServiceServer()
}

func RegisterBareksaNewsServiceServer(s grpc.ServiceRegistrar, srv BareksaNewsServiceServer) {
	s.RegisterService(&BareksaNewsService_ServiceDesc, srv)
}

func _BareksaNewsService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).AddTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_EditTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).EditTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/EditTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).EditTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Select)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).DeleteTag(ctx, req.(*Select))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).GetTags(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/AddTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).AddTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_EditTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).EditTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/EditTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).EditTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Select)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).DeleteTopic(ctx, req.(*Select))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).GetTopics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_AddNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(News)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).AddNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/AddNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).AddNews(ctx, req.(*News))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_EditNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(News)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).EditNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/EditNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).EditNews(ctx, req.(*News))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_DeleteNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Select)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).DeleteNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/DeleteNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).DeleteNews(ctx, req.(*Select))
	}
	return interceptor(ctx, in, info, handler)
}

func _BareksaNewsService_GetNewses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareksaNewsServiceServer).GetNewses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BareksaNewsService/GetNewses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareksaNewsServiceServer).GetNewses(ctx, req.(*Filters))
	}
	return interceptor(ctx, in, info, handler)
}

// BareksaNewsService_ServiceDesc is the grpc.ServiceDesc for BareksaNewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BareksaNewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.BareksaNewsService",
	HandlerType: (*BareksaNewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTag",
			Handler:    _BareksaNewsService_AddTag_Handler,
		},
		{
			MethodName: "EditTag",
			Handler:    _BareksaNewsService_EditTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _BareksaNewsService_DeleteTag_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _BareksaNewsService_GetTags_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _BareksaNewsService_AddTopic_Handler,
		},
		{
			MethodName: "EditTopic",
			Handler:    _BareksaNewsService_EditTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _BareksaNewsService_DeleteTopic_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _BareksaNewsService_GetTopics_Handler,
		},
		{
			MethodName: "AddNews",
			Handler:    _BareksaNewsService_AddNews_Handler,
		},
		{
			MethodName: "EditNews",
			Handler:    _BareksaNewsService_EditNews_Handler,
		},
		{
			MethodName: "DeleteNews",
			Handler:    _BareksaNewsService_DeleteNews_Handler,
		},
		{
			MethodName: "GetNewses",
			Handler:    _BareksaNewsService_GetNewses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tag.proto",
}
