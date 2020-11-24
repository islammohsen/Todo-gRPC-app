// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoResponse, error)
	GetAllTodos(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetAllTodosResponse, error)
	GetAllTodosStreaming(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (TodoService_GetAllTodosStreamingClient, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoResponse, error) {
	out := new(AddTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetAllTodos(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetAllTodosResponse, error) {
	out := new(GetAllTodosResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/GetAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetAllTodosStreaming(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (TodoService_GetAllTodosStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoService_serviceDesc.Streams[0], "/todo.TodoService/GetAllTodosStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceGetAllTodosStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_GetAllTodosStreamingClient interface {
	Recv() (*TodoItem, error)
	grpc.ClientStream
}

type todoServiceGetAllTodosStreamingClient struct {
	grpc.ClientStream
}

func (x *todoServiceGetAllTodosStreamingClient) Recv() (*TodoItem, error) {
	m := new(TodoItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	AddTodo(context.Context, *AddTodoRequest) (*AddTodoResponse, error)
	GetAllTodos(context.Context, *NoParams) (*GetAllTodosResponse, error)
	GetAllTodosStreaming(*NoParams, TodoService_GetAllTodosStreamingServer) error
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) AddTodo(context.Context, *AddTodoRequest) (*AddTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetAllTodos(context.Context, *NoParams) (*GetAllTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoServiceServer) GetAllTodosStreaming(*NoParams, TodoService_GetAllTodosStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllTodosStreaming not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodo(ctx, req.(*AddTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetAllTodos(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetAllTodosStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).GetAllTodosStreaming(m, &todoServiceGetAllTodosStreamingServer{stream})
}

type TodoService_GetAllTodosStreamingServer interface {
	Send(*TodoItem) error
	grpc.ServerStream
}

type todoServiceGetAllTodosStreamingServer struct {
	grpc.ServerStream
}

func (x *todoServiceGetAllTodosStreamingServer) Send(m *TodoItem) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _TodoService_AddTodo_Handler,
		},
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoService_GetAllTodos_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllTodosStreaming",
			Handler:       _TodoService_GetAllTodosStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}
