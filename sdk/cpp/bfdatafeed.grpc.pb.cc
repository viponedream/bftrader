// Generated by the gRPC protobuf plugin.
// If you make any local change, they will be lost.
// source: bfdatafeed.proto

#include "bfdatafeed.pb.h"
#include "bfdatafeed.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace bftrader {
namespace bfdatafeed {

static const char* BfDatafeedService_method_names[] = {
  "/bftrader.bfdatafeed.BfDatafeedService/Ping",
  "/bftrader.bfdatafeed.BfDatafeedService/InsertTick",
  "/bftrader.bfdatafeed.BfDatafeedService/GetTick",
  "/bftrader.bfdatafeed.BfDatafeedService/DeleteTick",
  "/bftrader.bfdatafeed.BfDatafeedService/InsertBar",
  "/bftrader.bfdatafeed.BfDatafeedService/GetBar",
  "/bftrader.bfdatafeed.BfDatafeedService/DeleteBar",
  "/bftrader.bfdatafeed.BfDatafeedService/InsertContract",
  "/bftrader.bfdatafeed.BfDatafeedService/GetContract",
  "/bftrader.bfdatafeed.BfDatafeedService/DeleteContract",
};

std::unique_ptr< BfDatafeedService::Stub> BfDatafeedService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< BfDatafeedService::Stub> stub(new BfDatafeedService::Stub(channel));
  return stub;
}

BfDatafeedService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_Ping_(BfDatafeedService_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_InsertTick_(BfDatafeedService_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetTick_(BfDatafeedService_method_names[2], ::grpc::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_DeleteTick_(BfDatafeedService_method_names[3], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_InsertBar_(BfDatafeedService_method_names[4], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetBar_(BfDatafeedService_method_names[5], ::grpc::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_DeleteBar_(BfDatafeedService_method_names[6], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_InsertContract_(BfDatafeedService_method_names[7], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetContract_(BfDatafeedService_method_names[8], ::grpc::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_DeleteContract_(BfDatafeedService_method_names[9], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status BfDatafeedService::Stub::Ping(::grpc::ClientContext* context, const ::bftrader::BfPingData& request, ::bftrader::BfPingData* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_Ping_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfPingData>* BfDatafeedService::Stub::AsyncPingRaw(::grpc::ClientContext* context, const ::bftrader::BfPingData& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfPingData>(channel_.get(), cq, rpcmethod_Ping_, context, request);
}

::grpc::Status BfDatafeedService::Stub::InsertTick(::grpc::ClientContext* context, const ::bftrader::BfTickData& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_InsertTick_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncInsertTickRaw(::grpc::ClientContext* context, const ::bftrader::BfTickData& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_InsertTick_, context, request);
}

::grpc::ClientReader< ::bftrader::BfTickData>* BfDatafeedService::Stub::GetTickRaw(::grpc::ClientContext* context, const ::bftrader::BfGetTickReq& request) {
  return new ::grpc::ClientReader< ::bftrader::BfTickData>(channel_.get(), rpcmethod_GetTick_, context, request);
}

::grpc::ClientAsyncReader< ::bftrader::BfTickData>* BfDatafeedService::Stub::AsyncGetTickRaw(::grpc::ClientContext* context, const ::bftrader::BfGetTickReq& request, ::grpc::CompletionQueue* cq, void* tag) {
  return new ::grpc::ClientAsyncReader< ::bftrader::BfTickData>(channel_.get(), cq, rpcmethod_GetTick_, context, request, tag);
}

::grpc::Status BfDatafeedService::Stub::DeleteTick(::grpc::ClientContext* context, const ::bftrader::BfDeleteTickReq& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_DeleteTick_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncDeleteTickRaw(::grpc::ClientContext* context, const ::bftrader::BfDeleteTickReq& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_DeleteTick_, context, request);
}

::grpc::Status BfDatafeedService::Stub::InsertBar(::grpc::ClientContext* context, const ::bftrader::BfBarData& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_InsertBar_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncInsertBarRaw(::grpc::ClientContext* context, const ::bftrader::BfBarData& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_InsertBar_, context, request);
}

::grpc::ClientReader< ::bftrader::BfBarData>* BfDatafeedService::Stub::GetBarRaw(::grpc::ClientContext* context, const ::bftrader::BfGetBarReq& request) {
  return new ::grpc::ClientReader< ::bftrader::BfBarData>(channel_.get(), rpcmethod_GetBar_, context, request);
}

::grpc::ClientAsyncReader< ::bftrader::BfBarData>* BfDatafeedService::Stub::AsyncGetBarRaw(::grpc::ClientContext* context, const ::bftrader::BfGetBarReq& request, ::grpc::CompletionQueue* cq, void* tag) {
  return new ::grpc::ClientAsyncReader< ::bftrader::BfBarData>(channel_.get(), cq, rpcmethod_GetBar_, context, request, tag);
}

::grpc::Status BfDatafeedService::Stub::DeleteBar(::grpc::ClientContext* context, const ::bftrader::BfDeleteBarReq& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_DeleteBar_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncDeleteBarRaw(::grpc::ClientContext* context, const ::bftrader::BfDeleteBarReq& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_DeleteBar_, context, request);
}

::grpc::Status BfDatafeedService::Stub::InsertContract(::grpc::ClientContext* context, const ::bftrader::BfContractData& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_InsertContract_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncInsertContractRaw(::grpc::ClientContext* context, const ::bftrader::BfContractData& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_InsertContract_, context, request);
}

::grpc::ClientReader< ::bftrader::BfContractData>* BfDatafeedService::Stub::GetContractRaw(::grpc::ClientContext* context, const ::bftrader::BfDatafeedGetContractReq& request) {
  return new ::grpc::ClientReader< ::bftrader::BfContractData>(channel_.get(), rpcmethod_GetContract_, context, request);
}

::grpc::ClientAsyncReader< ::bftrader::BfContractData>* BfDatafeedService::Stub::AsyncGetContractRaw(::grpc::ClientContext* context, const ::bftrader::BfDatafeedGetContractReq& request, ::grpc::CompletionQueue* cq, void* tag) {
  return new ::grpc::ClientAsyncReader< ::bftrader::BfContractData>(channel_.get(), cq, rpcmethod_GetContract_, context, request, tag);
}

::grpc::Status BfDatafeedService::Stub::DeleteContract(::grpc::ClientContext* context, const ::bftrader::BfDatafeedDeleteContractReq& request, ::bftrader::BfVoid* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_DeleteContract_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>* BfDatafeedService::Stub::AsyncDeleteContractRaw(::grpc::ClientContext* context, const ::bftrader::BfDatafeedDeleteContractReq& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::bftrader::BfVoid>(channel_.get(), cq, rpcmethod_DeleteContract_, context, request);
}

BfDatafeedService::Service::Service() {
  (void)BfDatafeedService_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfPingData, ::bftrader::BfPingData>(
          std::mem_fn(&BfDatafeedService::Service::Ping), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfTickData, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::InsertTick), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[2],
      ::grpc::RpcMethod::SERVER_STREAMING,
      new ::grpc::ServerStreamingHandler< BfDatafeedService::Service, ::bftrader::BfGetTickReq, ::bftrader::BfTickData>(
          std::mem_fn(&BfDatafeedService::Service::GetTick), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[3],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfDeleteTickReq, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::DeleteTick), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[4],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfBarData, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::InsertBar), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[5],
      ::grpc::RpcMethod::SERVER_STREAMING,
      new ::grpc::ServerStreamingHandler< BfDatafeedService::Service, ::bftrader::BfGetBarReq, ::bftrader::BfBarData>(
          std::mem_fn(&BfDatafeedService::Service::GetBar), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[6],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfDeleteBarReq, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::DeleteBar), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[7],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfContractData, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::InsertContract), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[8],
      ::grpc::RpcMethod::SERVER_STREAMING,
      new ::grpc::ServerStreamingHandler< BfDatafeedService::Service, ::bftrader::BfDatafeedGetContractReq, ::bftrader::BfContractData>(
          std::mem_fn(&BfDatafeedService::Service::GetContract), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      BfDatafeedService_method_names[9],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< BfDatafeedService::Service, ::bftrader::BfDatafeedDeleteContractReq, ::bftrader::BfVoid>(
          std::mem_fn(&BfDatafeedService::Service::DeleteContract), this)));
}

BfDatafeedService::Service::~Service() {
}

::grpc::Status BfDatafeedService::Service::Ping(::grpc::ServerContext* context, const ::bftrader::BfPingData* request, ::bftrader::BfPingData* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::InsertTick(::grpc::ServerContext* context, const ::bftrader::BfTickData* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::GetTick(::grpc::ServerContext* context, const ::bftrader::BfGetTickReq* request, ::grpc::ServerWriter< ::bftrader::BfTickData>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::DeleteTick(::grpc::ServerContext* context, const ::bftrader::BfDeleteTickReq* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::InsertBar(::grpc::ServerContext* context, const ::bftrader::BfBarData* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::GetBar(::grpc::ServerContext* context, const ::bftrader::BfGetBarReq* request, ::grpc::ServerWriter< ::bftrader::BfBarData>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::DeleteBar(::grpc::ServerContext* context, const ::bftrader::BfDeleteBarReq* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::InsertContract(::grpc::ServerContext* context, const ::bftrader::BfContractData* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::GetContract(::grpc::ServerContext* context, const ::bftrader::BfDatafeedGetContractReq* request, ::grpc::ServerWriter< ::bftrader::BfContractData>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status BfDatafeedService::Service::DeleteContract(::grpc::ServerContext* context, const ::bftrader::BfDatafeedDeleteContractReq* request, ::bftrader::BfVoid* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace bftrader
}  // namespace bfdatafeed

