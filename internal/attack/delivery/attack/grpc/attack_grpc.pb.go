// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// PokemonHandlerClient is the client API for PokemonHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonHandlerClient interface {
	GetPokemon(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Data, error)
}

type pokemonHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonHandlerClient(cc grpc.ClientConnInterface) PokemonHandlerClient {
	return &pokemonHandlerClient{cc}
}

func (c *pokemonHandlerClient) GetPokemon(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/attack.PokemonHandler/GetPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonHandlerServer is the server API for PokemonHandler service.
// All implementations must embed UnimplementedPokemonHandlerServer
// for forward compatibility
type PokemonHandlerServer interface {
	GetPokemon(context.Context, *Params) (*Data, error)
	mustEmbedUnimplementedPokemonHandlerServer()
}

// UnimplementedPokemonHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonHandlerServer struct {
}

func (UnimplementedPokemonHandlerServer) GetPokemon(context.Context, *Params) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedPokemonHandlerServer) mustEmbedUnimplementedPokemonHandlerServer() {}

// UnsafePokemonHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonHandlerServer will
// result in compilation errors.
type UnsafePokemonHandlerServer interface {
	mustEmbedUnimplementedPokemonHandlerServer()
}

func RegisterPokemonHandlerServer(s grpc.ServiceRegistrar, srv PokemonHandlerServer) {
	s.RegisterService(&PokemonHandler_ServiceDesc, srv)
}

func _PokemonHandler_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonHandlerServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attack.PokemonHandler/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonHandlerServer).GetPokemon(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

// PokemonHandler_ServiceDesc is the grpc.ServiceDesc for PokemonHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokemonHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attack.PokemonHandler",
	HandlerType: (*PokemonHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPokemon",
			Handler:    _PokemonHandler_GetPokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attack.proto",
}
