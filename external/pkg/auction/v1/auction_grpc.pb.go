// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: auction/v1/auction.proto

package auction

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
	Auction_UpBalance_FullMethodName           = "/Auction/UpBalance"
	Auction_CreateLot_FullMethodName           = "/Auction/CreateLot"
	Auction_FetchLot_FullMethodName            = "/Auction/FetchLot"
	Auction_ListLots_FullMethodName            = "/Auction/ListLots"
	Auction_CreateBid_FullMethodName           = "/Auction/CreateBid"
	Auction_CloseAuction_FullMethodName        = "/Auction/CloseAuction"
	Auction_ProcessTransactions_FullMethodName = "/Auction/ProcessTransactions"
)

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	UpBalance(ctx context.Context, in *UpBalanceRequest, opts ...grpc.CallOption) (*UpBalanceResponse, error)
	CreateLot(ctx context.Context, in *CreateLotRequest, opts ...grpc.CallOption) (*CreateLotResponse, error)
	FetchLot(ctx context.Context, in *FetchLotRequest, opts ...grpc.CallOption) (*FetchLotResponse, error)
	ListLots(ctx context.Context, in *ListLotsRequest, opts ...grpc.CallOption) (*ListLotsResponse, error)
	CreateBid(ctx context.Context, in *CreateBidRequest, opts ...grpc.CallOption) (*CreateBidResponse, error)
	CloseAuction(ctx context.Context, in *CloseAuctionRequest, opts ...grpc.CallOption) (*CloseAuctionResponse, error)
	ProcessTransactions(ctx context.Context, in *ProcessTransactionsRequest, opts ...grpc.CallOption) (*ProcessTransactionsResponse, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) UpBalance(ctx context.Context, in *UpBalanceRequest, opts ...grpc.CallOption) (*UpBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpBalanceResponse)
	err := c.cc.Invoke(ctx, Auction_UpBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) CreateLot(ctx context.Context, in *CreateLotRequest, opts ...grpc.CallOption) (*CreateLotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLotResponse)
	err := c.cc.Invoke(ctx, Auction_CreateLot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) FetchLot(ctx context.Context, in *FetchLotRequest, opts ...grpc.CallOption) (*FetchLotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchLotResponse)
	err := c.cc.Invoke(ctx, Auction_FetchLot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) ListLots(ctx context.Context, in *ListLotsRequest, opts ...grpc.CallOption) (*ListLotsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLotsResponse)
	err := c.cc.Invoke(ctx, Auction_ListLots_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) CreateBid(ctx context.Context, in *CreateBidRequest, opts ...grpc.CallOption) (*CreateBidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBidResponse)
	err := c.cc.Invoke(ctx, Auction_CreateBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) CloseAuction(ctx context.Context, in *CloseAuctionRequest, opts ...grpc.CallOption) (*CloseAuctionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseAuctionResponse)
	err := c.cc.Invoke(ctx, Auction_CloseAuction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) ProcessTransactions(ctx context.Context, in *ProcessTransactionsRequest, opts ...grpc.CallOption) (*ProcessTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessTransactionsResponse)
	err := c.cc.Invoke(ctx, Auction_ProcessTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility.
type AuctionServer interface {
	UpBalance(context.Context, *UpBalanceRequest) (*UpBalanceResponse, error)
	CreateLot(context.Context, *CreateLotRequest) (*CreateLotResponse, error)
	FetchLot(context.Context, *FetchLotRequest) (*FetchLotResponse, error)
	ListLots(context.Context, *ListLotsRequest) (*ListLotsResponse, error)
	CreateBid(context.Context, *CreateBidRequest) (*CreateBidResponse, error)
	CloseAuction(context.Context, *CloseAuctionRequest) (*CloseAuctionResponse, error)
	ProcessTransactions(context.Context, *ProcessTransactionsRequest) (*ProcessTransactionsResponse, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionServer struct{}

func (UnimplementedAuctionServer) UpBalance(context.Context, *UpBalanceRequest) (*UpBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpBalance not implemented")
}
func (UnimplementedAuctionServer) CreateLot(context.Context, *CreateLotRequest) (*CreateLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLot not implemented")
}
func (UnimplementedAuctionServer) FetchLot(context.Context, *FetchLotRequest) (*FetchLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLot not implemented")
}
func (UnimplementedAuctionServer) ListLots(context.Context, *ListLotsRequest) (*ListLotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLots not implemented")
}
func (UnimplementedAuctionServer) CreateBid(context.Context, *CreateBidRequest) (*CreateBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBid not implemented")
}
func (UnimplementedAuctionServer) CloseAuction(context.Context, *CloseAuctionRequest) (*CloseAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseAuction not implemented")
}
func (UnimplementedAuctionServer) ProcessTransactions(context.Context, *ProcessTransactionsRequest) (*ProcessTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessTransactions not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}
func (UnimplementedAuctionServer) testEmbeddedByValue()                 {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	// If the following call pancis, it indicates UnimplementedAuctionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_UpBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).UpBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_UpBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).UpBalance(ctx, req.(*UpBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_CreateLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).CreateLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_CreateLot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).CreateLot(ctx, req.(*CreateLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_FetchLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).FetchLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_FetchLot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).FetchLot(ctx, req.(*FetchLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_ListLots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).ListLots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_ListLots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).ListLots(ctx, req.(*ListLotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_CreateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).CreateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_CreateBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).CreateBid(ctx, req.(*CreateBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_CloseAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).CloseAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_CloseAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).CloseAuction(ctx, req.(*CloseAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_ProcessTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).ProcessTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_ProcessTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).ProcessTransactions(ctx, req.(*ProcessTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpBalance",
			Handler:    _Auction_UpBalance_Handler,
		},
		{
			MethodName: "CreateLot",
			Handler:    _Auction_CreateLot_Handler,
		},
		{
			MethodName: "FetchLot",
			Handler:    _Auction_FetchLot_Handler,
		},
		{
			MethodName: "ListLots",
			Handler:    _Auction_ListLots_Handler,
		},
		{
			MethodName: "CreateBid",
			Handler:    _Auction_CreateBid_Handler,
		},
		{
			MethodName: "CloseAuction",
			Handler:    _Auction_CloseAuction_Handler,
		},
		{
			MethodName: "ProcessTransactions",
			Handler:    _Auction_ProcessTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/v1/auction.proto",
}
