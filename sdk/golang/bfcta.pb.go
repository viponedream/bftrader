// Code generated by protoc-gen-go.
// source: bfcta.proto
// DO NOT EDIT!

/*
Package bftrader_bfcta is a generated protocol buffer package.

It is generated from these files:
	bfcta.proto

It has these top-level messages:
*/
package bftrader_bfcta

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

// Client API for BfCtaService service

type BfCtaServiceClient interface {
	// 建立连接
	Connect(ctx context.Context, in *bftrader.BfConnectReq, opts ...grpc.CallOption) (*bftrader.BfConnectResp, error)
	// 获取策略信息，如position，用kv最方便
	GetRobotInfo(ctx context.Context, in *bftrader.BfKvData, opts ...grpc.CallOption) (*bftrader.BfKvData, error)
	// 发单
	SendOrder(ctx context.Context, in *bftrader.BfSendOrderReq, opts ...grpc.CallOption) (*bftrader.BfSendOrderResp, error)
	// 撤单
	CancelOrder(ctx context.Context, in *bftrader.BfCancelOrderReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 关闭连接
	Close(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
}

type bfCtaServiceClient struct {
	cc *grpc.ClientConn
}

func NewBfCtaServiceClient(cc *grpc.ClientConn) BfCtaServiceClient {
	return &bfCtaServiceClient{cc}
}

func (c *bfCtaServiceClient) Connect(ctx context.Context, in *bftrader.BfConnectReq, opts ...grpc.CallOption) (*bftrader.BfConnectResp, error) {
	out := new(bftrader.BfConnectResp)
	err := grpc.Invoke(ctx, "/bftrader.bfcta.BfCtaService/Connect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) GetRobotInfo(ctx context.Context, in *bftrader.BfKvData, opts ...grpc.CallOption) (*bftrader.BfKvData, error) {
	out := new(bftrader.BfKvData)
	err := grpc.Invoke(ctx, "/bftrader.bfcta.BfCtaService/GetRobotInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) SendOrder(ctx context.Context, in *bftrader.BfSendOrderReq, opts ...grpc.CallOption) (*bftrader.BfSendOrderResp, error) {
	out := new(bftrader.BfSendOrderResp)
	err := grpc.Invoke(ctx, "/bftrader.bfcta.BfCtaService/SendOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) CancelOrder(ctx context.Context, in *bftrader.BfCancelOrderReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfcta.BfCtaService/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) Close(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfcta.BfCtaService/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BfCtaService service

type BfCtaServiceServer interface {
	// 建立连接
	Connect(context.Context, *bftrader.BfConnectReq) (*bftrader.BfConnectResp, error)
	// 获取策略信息，如position，用kv最方便
	GetRobotInfo(context.Context, *bftrader.BfKvData) (*bftrader.BfKvData, error)
	// 发单
	SendOrder(context.Context, *bftrader.BfSendOrderReq) (*bftrader.BfSendOrderResp, error)
	// 撤单
	CancelOrder(context.Context, *bftrader.BfCancelOrderReq) (*bftrader.BfVoid, error)
	// 关闭连接
	Close(context.Context, *bftrader.BfVoid) (*bftrader.BfVoid, error)
}

func RegisterBfCtaServiceServer(s *grpc.Server, srv BfCtaServiceServer) {
	s.RegisterService(&_BfCtaService_serviceDesc, srv)
}

func _BfCtaService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfCtaServiceServer).Connect(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfCtaService_GetRobotInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfKvData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfCtaServiceServer).GetRobotInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfCtaService_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfSendOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfCtaServiceServer).SendOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfCtaService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfCancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfCtaServiceServer).CancelOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfCtaService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfCtaServiceServer).Close(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BfCtaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftrader.bfcta.BfCtaService",
	HandlerType: (*BfCtaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _BfCtaService_Connect_Handler,
		},
		{
			MethodName: "GetRobotInfo",
			Handler:    _BfCtaService_GetRobotInfo_Handler,
		},
		{
			MethodName: "SendOrder",
			Handler:    _BfCtaService_SendOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _BfCtaService_CancelOrder_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _BfCtaService_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4a, 0x4b, 0x2e,
	0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x4a, 0x2b, 0x29, 0x4a, 0x4c, 0x49,
	0x2d, 0xd2, 0x03, 0x8b, 0x4a, 0x21, 0xf8, 0x60, 0x79, 0xa3, 0x1d, 0x4c, 0x5c, 0x3c, 0x4e, 0x69,
	0xce, 0x25, 0x89, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x36, 0x5c, 0xec, 0xce, 0xf9,
	0x79, 0x79, 0xa9, 0xc9, 0x25, 0x42, 0x62, 0x7a, 0x70, 0xc5, 0x40, 0x25, 0x10, 0xc1, 0xa0, 0xd4,
	0x42, 0x29, 0x71, 0xac, 0xe2, 0xc5, 0x05, 0x4a, 0x0c, 0x42, 0x16, 0x5c, 0x3c, 0xee, 0xa9, 0x25,
	0x41, 0xf9, 0x49, 0xf9, 0x25, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x42, 0xc8, 0x4a, 0xbd, 0xcb, 0x5c,
	0x12, 0x81, 0x4e, 0xc0, 0x22, 0x06, 0xd4, 0xe9, 0xc4, 0xc5, 0x19, 0x9c, 0x9a, 0x97, 0xe2, 0x5f,
	0x04, 0x14, 0x17, 0x92, 0x40, 0x56, 0x02, 0x17, 0x06, 0xd9, 0x2d, 0x89, 0x43, 0x06, 0x6c, 0xbb,
	0x2d, 0x17, 0xb7, 0x73, 0x62, 0x5e, 0x72, 0x6a, 0x0e, 0xc4, 0x14, 0x29, 0x14, 0x77, 0x22, 0x24,
	0x40, 0xe6, 0x08, 0x20, 0xcb, 0x85, 0xe5, 0x67, 0xa6, 0x00, 0xb5, 0xeb, 0x72, 0xb1, 0x3a, 0xe7,
	0xe4, 0x17, 0xa7, 0x0a, 0x61, 0x48, 0x62, 0x53, 0x9e, 0xc4, 0x06, 0x0e, 0x41, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x1f, 0x20, 0x0e, 0x70, 0x01, 0x00, 0x00,
}
