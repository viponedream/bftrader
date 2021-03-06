// Code generated by protoc-gen-go.
// source: bfdatafeed.proto
// DO NOT EDIT!

/*
Package bftrader_bfdatafeed is a generated protocol buffer package.

It is generated from these files:
	bfdatafeed.proto

It has these top-level messages:
*/
package bftrader_bfdatafeed

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bftrader "."

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for BfDatafeedService service

type BfDatafeedServiceClient interface {
	// 活跃检测
	Ping(ctx context.Context, in *bftrader.BfPingData, opts ...grpc.CallOption) (*bftrader.BfPingData, error)
	// 保存tick
	InsertTick(ctx context.Context, in *bftrader.BfTickData, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 获取tick
	GetTick(ctx context.Context, in *bftrader.BfGetTickReq, opts ...grpc.CallOption) (BfDatafeedService_GetTickClient, error)
	// 删除tick
	DeleteTick(ctx context.Context, in *bftrader.BfDeleteTickReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 保存bar
	InsertBar(ctx context.Context, in *bftrader.BfBarData, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 获取bar
	GetBar(ctx context.Context, in *bftrader.BfGetBarReq, opts ...grpc.CallOption) (BfDatafeedService_GetBarClient, error)
	// 删除bar
	DeleteBar(ctx context.Context, in *bftrader.BfDeleteBarReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 保存contract
	InsertContract(ctx context.Context, in *bftrader.BfContractData, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 获取contract
	GetContract(ctx context.Context, in *bftrader.BfDatafeedGetContractReq, opts ...grpc.CallOption) (BfDatafeedService_GetContractClient, error)
	// 删除contract
	DeleteContract(ctx context.Context, in *bftrader.BfDatafeedDeleteContractReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
}

type bfDatafeedServiceClient struct {
	cc *grpc.ClientConn
}

func NewBfDatafeedServiceClient(cc *grpc.ClientConn) BfDatafeedServiceClient {
	return &bfDatafeedServiceClient{cc}
}

func (c *bfDatafeedServiceClient) Ping(ctx context.Context, in *bftrader.BfPingData, opts ...grpc.CallOption) (*bftrader.BfPingData, error) {
	out := new(bftrader.BfPingData)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertTick(ctx context.Context, in *bftrader.BfTickData, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/InsertTick", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) GetTick(ctx context.Context, in *bftrader.BfGetTickReq, opts ...grpc.CallOption) (BfDatafeedService_GetTickClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[0], c.cc, "/bftrader.bfdatafeed.BfDatafeedService/GetTick", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetTickClient interface {
	Recv() (*bftrader.BfTickData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetTickClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetTickClient) Recv() (*bftrader.BfTickData, error) {
	m := new(bftrader.BfTickData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfDatafeedServiceClient) DeleteTick(ctx context.Context, in *bftrader.BfDeleteTickReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/DeleteTick", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertBar(ctx context.Context, in *bftrader.BfBarData, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/InsertBar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) GetBar(ctx context.Context, in *bftrader.BfGetBarReq, opts ...grpc.CallOption) (BfDatafeedService_GetBarClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[1], c.cc, "/bftrader.bfdatafeed.BfDatafeedService/GetBar", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetBarClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetBarClient interface {
	Recv() (*bftrader.BfBarData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetBarClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetBarClient) Recv() (*bftrader.BfBarData, error) {
	m := new(bftrader.BfBarData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfDatafeedServiceClient) DeleteBar(ctx context.Context, in *bftrader.BfDeleteBarReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/DeleteBar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) InsertContract(ctx context.Context, in *bftrader.BfContractData, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/InsertContract", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfDatafeedServiceClient) GetContract(ctx context.Context, in *bftrader.BfDatafeedGetContractReq, opts ...grpc.CallOption) (BfDatafeedService_GetContractClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfDatafeedService_serviceDesc.Streams[2], c.cc, "/bftrader.bfdatafeed.BfDatafeedService/GetContract", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfDatafeedServiceGetContractClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfDatafeedService_GetContractClient interface {
	Recv() (*bftrader.BfContractData, error)
	grpc.ClientStream
}

type bfDatafeedServiceGetContractClient struct {
	grpc.ClientStream
}

func (x *bfDatafeedServiceGetContractClient) Recv() (*bftrader.BfContractData, error) {
	m := new(bftrader.BfContractData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfDatafeedServiceClient) DeleteContract(ctx context.Context, in *bftrader.BfDatafeedDeleteContractReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfdatafeed.BfDatafeedService/DeleteContract", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BfDatafeedService service

type BfDatafeedServiceServer interface {
	// 活跃检测
	Ping(context.Context, *bftrader.BfPingData) (*bftrader.BfPingData, error)
	// 保存tick
	InsertTick(context.Context, *bftrader.BfTickData) (*bftrader.BfVoid, error)
	// 获取tick
	GetTick(*bftrader.BfGetTickReq, BfDatafeedService_GetTickServer) error
	// 删除tick
	DeleteTick(context.Context, *bftrader.BfDeleteTickReq) (*bftrader.BfVoid, error)
	// 保存bar
	InsertBar(context.Context, *bftrader.BfBarData) (*bftrader.BfVoid, error)
	// 获取bar
	GetBar(*bftrader.BfGetBarReq, BfDatafeedService_GetBarServer) error
	// 删除bar
	DeleteBar(context.Context, *bftrader.BfDeleteBarReq) (*bftrader.BfVoid, error)
	// 保存contract
	InsertContract(context.Context, *bftrader.BfContractData) (*bftrader.BfVoid, error)
	// 获取contract
	GetContract(*bftrader.BfDatafeedGetContractReq, BfDatafeedService_GetContractServer) error
	// 删除contract
	DeleteContract(context.Context, *bftrader.BfDatafeedDeleteContractReq) (*bftrader.BfVoid, error)
}

func RegisterBfDatafeedServiceServer(s *grpc.Server, srv BfDatafeedServiceServer) {
	s.RegisterService(&_BfDatafeedService_serviceDesc, srv)
}

func _BfDatafeedService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfPingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertTick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfTickData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertTick(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_GetTick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(bftrader.BfGetTickReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetTick(m, &bfDatafeedServiceGetTickServer{stream})
}

type BfDatafeedService_GetTickServer interface {
	Send(*bftrader.BfTickData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetTickServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetTickServer) Send(m *bftrader.BfTickData) error {
	return x.ServerStream.SendMsg(m)
}

func _BfDatafeedService_DeleteTick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfDeleteTickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).DeleteTick(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfBarData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertBar(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_GetBar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(bftrader.BfGetBarReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetBar(m, &bfDatafeedServiceGetBarServer{stream})
}

type BfDatafeedService_GetBarServer interface {
	Send(*bftrader.BfBarData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetBarServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetBarServer) Send(m *bftrader.BfBarData) error {
	return x.ServerStream.SendMsg(m)
}

func _BfDatafeedService_DeleteBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfDeleteBarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).DeleteBar(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_InsertContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfContractData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).InsertContract(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfDatafeedService_GetContract_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(bftrader.BfDatafeedGetContractReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfDatafeedServiceServer).GetContract(m, &bfDatafeedServiceGetContractServer{stream})
}

type BfDatafeedService_GetContractServer interface {
	Send(*bftrader.BfContractData) error
	grpc.ServerStream
}

type bfDatafeedServiceGetContractServer struct {
	grpc.ServerStream
}

func (x *bfDatafeedServiceGetContractServer) Send(m *bftrader.BfContractData) error {
	return x.ServerStream.SendMsg(m)
}

func _BfDatafeedService_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfDatafeedDeleteContractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfDatafeedServiceServer).DeleteContract(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BfDatafeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftrader.bfdatafeed.BfDatafeedService",
	HandlerType: (*BfDatafeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BfDatafeedService_Ping_Handler,
		},
		{
			MethodName: "InsertTick",
			Handler:    _BfDatafeedService_InsertTick_Handler,
		},
		{
			MethodName: "DeleteTick",
			Handler:    _BfDatafeedService_DeleteTick_Handler,
		},
		{
			MethodName: "InsertBar",
			Handler:    _BfDatafeedService_InsertBar_Handler,
		},
		{
			MethodName: "DeleteBar",
			Handler:    _BfDatafeedService_DeleteBar_Handler,
		},
		{
			MethodName: "InsertContract",
			Handler:    _BfDatafeedService_InsertContract_Handler,
		},
		{
			MethodName: "DeleteContract",
			Handler:    _BfDatafeedService_DeleteContract_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTick",
			Handler:       _BfDatafeedService_GetTick_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBar",
			Handler:       _BfDatafeedService_GetBar_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetContract",
			Handler:       _BfDatafeedService_GetContract_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x57, 0xd0, 0x95, 0x1d, 0xa1, 0xac, 0x59, 0x15, 0xed, 0xb1, 0xe0, 0xb5, 0x88, 0x2e,
	0xa2, 0x08, 0x1e, 0xea, 0x82, 0x88, 0x07, 0x45, 0xc5, 0x7b, 0x3f, 0x26, 0x52, 0x94, 0x56, 0x63,
	0xf0, 0x17, 0xf8, 0xc3, 0xcd, 0x24, 0x35, 0x4d, 0x34, 0xb9, 0xbe, 0xf3, 0x3e, 0x33, 0x4f, 0x68,
	0x61, 0x5e, 0xf1, 0xa6, 0x94, 0x25, 0x47, 0x6c, 0xf2, 0x77, 0xd1, 0xcb, 0x9e, 0x2d, 0x2a, 0x2e,
	0x45, 0xd9, 0xa0, 0xc8, 0xc7, 0x51, 0x9a, 0xd8, 0x50, 0x97, 0x8e, 0xbf, 0x37, 0x60, 0xbb, 0xe0,
	0xab, 0x61, 0xfc, 0x88, 0xe2, 0xab, 0xad, 0x91, 0x2d, 0x61, 0xfd, 0xbe, 0xed, 0x5e, 0xd8, 0x4e,
	0x6e, 0xeb, 0x05, 0xa7, 0x84, 0x8a, 0x69, 0x30, 0xcd, 0x26, 0xec, 0x14, 0xe0, 0xa6, 0xfb, 0x44,
	0x21, 0x9f, 0xda, 0xfa, 0xd5, 0x67, 0x29, 0xd1, 0xec, 0xdc, 0x4d, 0x9f, 0xfb, 0xb6, 0x51, 0xdc,
	0x05, 0x6c, 0x5e, 0xa3, 0x81, 0xf6, 0xdc, 0xf1, 0x10, 0x3e, 0xe0, 0x47, 0x1a, 0x5c, 0x96, 0x4d,
	0x8e, 0xd6, 0x14, 0x0c, 0x2b, 0x7c, 0x43, 0x89, 0x9a, 0x3f, 0x70, 0x7b, 0x63, 0x4e, 0x2b, 0x42,
	0x97, 0x97, 0x30, 0x33, 0xc6, 0x45, 0x29, 0xd8, 0xc2, 0x2d, 0xa8, 0x20, 0xea, 0x7b, 0x06, 0x53,
	0xa5, 0x46, 0xc8, 0xee, 0x1f, 0x5d, 0x95, 0xd1, 0xa9, 0xd0, 0x26, 0x2d, 0x7b, 0x0e, 0x33, 0x23,
	0x45, 0xf0, 0xfe, 0x7f, 0xd7, 0x81, 0x0f, 0x1d, 0xbd, 0x84, 0xc4, 0xa8, 0x5e, 0xf5, 0x9d, 0x9a,
	0xd5, 0xd2, 0xe7, 0x7f, 0xd3, 0xa8, 0xf4, 0x1d, 0x6c, 0x29, 0x41, 0x0b, 0x67, 0xde, 0xf1, 0xe1,
	0xf3, 0x3b, 0x05, 0xd2, 0x88, 0x1e, 0xd0, 0x6f, 0xb9, 0x85, 0xc4, 0x48, 0xdb, 0x9d, 0x87, 0xa1,
	0x9d, 0x7e, 0x27, 0xf2, 0xba, 0x6a, 0xaa, 0xff, 0xc6, 0x93, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0xb3, 0x6f, 0x39, 0xc6, 0x02, 0x00, 0x00,
}
