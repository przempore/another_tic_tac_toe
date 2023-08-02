// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/tictactoe.proto

package tictactoe

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

// TicTacToeGrpcClient is the client API for TicTacToeGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicTacToeGrpcClient interface {
	InitialState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Board, error)
	Turn(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Player, error)
	PossibleActions(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Actions, error)
	Result(ctx context.Context, in *BoardWithAction, opts ...grpc.CallOption) (*Board, error)
	Winner(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Player, error)
	IsTerminal(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Terminal, error)
	Minimax(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Action, error)
}

type ticTacToeGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTicTacToeGrpcClient(cc grpc.ClientConnInterface) TicTacToeGrpcClient {
	return &ticTacToeGrpcClient{cc}
}

func (c *ticTacToeGrpcClient) InitialState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Board, error) {
	out := new(Board)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/InitialState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) Turn(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/Turn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) PossibleActions(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Actions, error) {
	out := new(Actions)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/PossibleActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) Result(ctx context.Context, in *BoardWithAction, opts ...grpc.CallOption) (*Board, error) {
	out := new(Board)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) Winner(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/Winner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) IsTerminal(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Terminal, error) {
	out := new(Terminal)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/IsTerminal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticTacToeGrpcClient) Minimax(ctx context.Context, in *Board, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/tictactoe.TicTacToeGrpc/Minimax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicTacToeGrpcServer is the server API for TicTacToeGrpc service.
// All implementations must embed UnimplementedTicTacToeGrpcServer
// for forward compatibility
type TicTacToeGrpcServer interface {
	InitialState(context.Context, *emptypb.Empty) (*Board, error)
	Turn(context.Context, *Board) (*Player, error)
	PossibleActions(context.Context, *Board) (*Actions, error)
	Result(context.Context, *BoardWithAction) (*Board, error)
	Winner(context.Context, *Board) (*Player, error)
	IsTerminal(context.Context, *Board) (*Terminal, error)
	Minimax(context.Context, *Board) (*Action, error)
	mustEmbedUnimplementedTicTacToeGrpcServer()
}

// UnimplementedTicTacToeGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedTicTacToeGrpcServer struct {
}

func (UnimplementedTicTacToeGrpcServer) InitialState(context.Context, *emptypb.Empty) (*Board, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialState not implemented")
}
func (UnimplementedTicTacToeGrpcServer) Turn(context.Context, *Board) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Turn not implemented")
}
func (UnimplementedTicTacToeGrpcServer) PossibleActions(context.Context, *Board) (*Actions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PossibleActions not implemented")
}
func (UnimplementedTicTacToeGrpcServer) Result(context.Context, *BoardWithAction) (*Board, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedTicTacToeGrpcServer) Winner(context.Context, *Board) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Winner not implemented")
}
func (UnimplementedTicTacToeGrpcServer) IsTerminal(context.Context, *Board) (*Terminal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTerminal not implemented")
}
func (UnimplementedTicTacToeGrpcServer) Minimax(context.Context, *Board) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Minimax not implemented")
}
func (UnimplementedTicTacToeGrpcServer) mustEmbedUnimplementedTicTacToeGrpcServer() {}

// UnsafeTicTacToeGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicTacToeGrpcServer will
// result in compilation errors.
type UnsafeTicTacToeGrpcServer interface {
	mustEmbedUnimplementedTicTacToeGrpcServer()
}

func RegisterTicTacToeGrpcServer(s grpc.ServiceRegistrar, srv TicTacToeGrpcServer) {
	s.RegisterService(&TicTacToeGrpc_ServiceDesc, srv)
}

func _TicTacToeGrpc_InitialState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).InitialState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/InitialState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).InitialState(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_Turn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Board)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).Turn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/Turn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).Turn(ctx, req.(*Board))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_PossibleActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Board)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).PossibleActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/PossibleActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).PossibleActions(ctx, req.(*Board))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardWithAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).Result(ctx, req.(*BoardWithAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_Winner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Board)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).Winner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/Winner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).Winner(ctx, req.(*Board))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_IsTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Board)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).IsTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/IsTerminal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).IsTerminal(ctx, req.(*Board))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicTacToeGrpc_Minimax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Board)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicTacToeGrpcServer).Minimax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.TicTacToeGrpc/Minimax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicTacToeGrpcServer).Minimax(ctx, req.(*Board))
	}
	return interceptor(ctx, in, info, handler)
}

// TicTacToeGrpc_ServiceDesc is the grpc.ServiceDesc for TicTacToeGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicTacToeGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tictactoe.TicTacToeGrpc",
	HandlerType: (*TicTacToeGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitialState",
			Handler:    _TicTacToeGrpc_InitialState_Handler,
		},
		{
			MethodName: "Turn",
			Handler:    _TicTacToeGrpc_Turn_Handler,
		},
		{
			MethodName: "PossibleActions",
			Handler:    _TicTacToeGrpc_PossibleActions_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _TicTacToeGrpc_Result_Handler,
		},
		{
			MethodName: "Winner",
			Handler:    _TicTacToeGrpc_Winner_Handler,
		},
		{
			MethodName: "IsTerminal",
			Handler:    _TicTacToeGrpc_IsTerminal_Handler,
		},
		{
			MethodName: "Minimax",
			Handler:    _TicTacToeGrpc_Minimax_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tictactoe.proto",
}
