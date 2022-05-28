// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bitcoindv22

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

// BitcoinClient is the client API for Bitcoin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BitcoinClient interface {
	GetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error)
}

type bitcoinClient struct {
	cc grpc.ClientConnInterface
}

func NewBitcoinClient(cc grpc.ClientConnInterface) BitcoinClient {
	return &bitcoinClient{cc}
}

func (c *bitcoinClient) GetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error) {
	out := new(GetBlockchainInfoResponse)
	err := c.cc.Invoke(ctx, "/barebitcoin.bitcoind.v22.Bitcoin/GetBlockchainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitcoinServer is the server API for Bitcoin service.
// All implementations should embed UnimplementedBitcoinServer
// for forward compatibility
type BitcoinServer interface {
	GetBlockchainInfo(context.Context, *GetBlockchainInfoRequest) (*GetBlockchainInfoResponse, error)
}

// UnimplementedBitcoinServer should be embedded to have forward compatible implementations.
type UnimplementedBitcoinServer struct {
}

func (UnimplementedBitcoinServer) GetBlockchainInfo(context.Context, *GetBlockchainInfoRequest) (*GetBlockchainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchainInfo not implemented")
}

// UnsafeBitcoinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BitcoinServer will
// result in compilation errors.
type UnsafeBitcoinServer interface {
	mustEmbedUnimplementedBitcoinServer()
}

func RegisterBitcoinServer(s grpc.ServiceRegistrar, srv BitcoinServer) {
	s.RegisterService(&Bitcoin_ServiceDesc, srv)
}

func _Bitcoin_GetBlockchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinServer).GetBlockchainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/barebitcoin.bitcoind.v22.Bitcoin/GetBlockchainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinServer).GetBlockchainInfo(ctx, req.(*GetBlockchainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bitcoin_ServiceDesc is the grpc.ServiceDesc for Bitcoin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bitcoin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "barebitcoin.bitcoind.v22.Bitcoin",
	HandlerType: (*BitcoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockchainInfo",
			Handler:    _Bitcoin_GetBlockchainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "barebitcoin/bitcoind/v22/bitcoin.proto",
}
