# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bfrobot.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import bftrader_pb2 as bftrader__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='bfrobot.proto',
  package='bftrader.bfrobot',
  syntax='proto3',
  serialized_pb=_b('\n\rbfrobot.proto\x12\x10\x62\x66trader.bfrobot\x1a\x0e\x62\x66trader.proto2\xf9\x02\n\x0e\x42\x66RobotService\x12\x36\n\x06OnPing\x12\x14.bftrader.BfPingData\x1a\x14.bftrader.BfPingData\"\x00\x12\x32\n\x06OnTick\x12\x14.bftrader.BfTickData\x1a\x10.bftrader.BfVoid\"\x00\x12\x34\n\x07OnTrade\x12\x15.bftrader.BfTradeData\x1a\x10.bftrader.BfVoid\"\x00\x12\x34\n\x07OnOrder\x12\x15.bftrader.BfOrderData\x1a\x10.bftrader.BfVoid\"\x00\x12.\n\x06OnInit\x12\x10.bftrader.BfVoid\x1a\x10.bftrader.BfVoid\"\x00\x12/\n\x07OnStart\x12\x10.bftrader.BfVoid\x1a\x10.bftrader.BfVoid\"\x00\x12.\n\x06OnStop\x12\x10.bftrader.BfVoid\x1a\x10.bftrader.BfVoid\"\x00\x62\x06proto3')
  ,
  dependencies=[bftrader__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)





import abc
from grpc.beta import implementations as beta_implementations
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

class BetaBfRobotServiceServicer(object):
  """<fill me in later!>"""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def OnPing(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnTick(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnTrade(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnOrder(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnInit(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnStart(self, request, context):
    raise NotImplementedError()
  @abc.abstractmethod
  def OnStop(self, request, context):
    raise NotImplementedError()

class BetaBfRobotServiceStub(object):
  """The interface to which stubs will conform."""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def OnPing(self, request, timeout):
    raise NotImplementedError()
  OnPing.future = None
  @abc.abstractmethod
  def OnTick(self, request, timeout):
    raise NotImplementedError()
  OnTick.future = None
  @abc.abstractmethod
  def OnTrade(self, request, timeout):
    raise NotImplementedError()
  OnTrade.future = None
  @abc.abstractmethod
  def OnOrder(self, request, timeout):
    raise NotImplementedError()
  OnOrder.future = None
  @abc.abstractmethod
  def OnInit(self, request, timeout):
    raise NotImplementedError()
  OnInit.future = None
  @abc.abstractmethod
  def OnStart(self, request, timeout):
    raise NotImplementedError()
  OnStart.future = None
  @abc.abstractmethod
  def OnStop(self, request, timeout):
    raise NotImplementedError()
  OnStop.future = None

def beta_create_BfRobotService_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  request_deserializers = {
    ('bftrader.bfrobot.BfRobotService', 'OnInit'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnOrder'): bftrader_pb2.BfOrderData.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnPing'): bftrader_pb2.BfPingData.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnStart'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnStop'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnTick'): bftrader_pb2.BfTickData.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnTrade'): bftrader_pb2.BfTradeData.FromString,
  }
  response_serializers = {
    ('bftrader.bfrobot.BfRobotService', 'OnInit'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnOrder'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnPing'): bftrader_pb2.BfPingData.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnStart'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnStop'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnTick'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnTrade'): bftrader_pb2.BfVoid.SerializeToString,
  }
  method_implementations = {
    ('bftrader.bfrobot.BfRobotService', 'OnInit'): face_utilities.unary_unary_inline(servicer.OnInit),
    ('bftrader.bfrobot.BfRobotService', 'OnOrder'): face_utilities.unary_unary_inline(servicer.OnOrder),
    ('bftrader.bfrobot.BfRobotService', 'OnPing'): face_utilities.unary_unary_inline(servicer.OnPing),
    ('bftrader.bfrobot.BfRobotService', 'OnStart'): face_utilities.unary_unary_inline(servicer.OnStart),
    ('bftrader.bfrobot.BfRobotService', 'OnStop'): face_utilities.unary_unary_inline(servicer.OnStop),
    ('bftrader.bfrobot.BfRobotService', 'OnTick'): face_utilities.unary_unary_inline(servicer.OnTick),
    ('bftrader.bfrobot.BfRobotService', 'OnTrade'): face_utilities.unary_unary_inline(servicer.OnTrade),
  }
  server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
  return beta_implementations.server(method_implementations, options=server_options)

def beta_create_BfRobotService_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  import bftrader_pb2
  request_serializers = {
    ('bftrader.bfrobot.BfRobotService', 'OnInit'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnOrder'): bftrader_pb2.BfOrderData.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnPing'): bftrader_pb2.BfPingData.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnStart'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnStop'): bftrader_pb2.BfVoid.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnTick'): bftrader_pb2.BfTickData.SerializeToString,
    ('bftrader.bfrobot.BfRobotService', 'OnTrade'): bftrader_pb2.BfTradeData.SerializeToString,
  }
  response_deserializers = {
    ('bftrader.bfrobot.BfRobotService', 'OnInit'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnOrder'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnPing'): bftrader_pb2.BfPingData.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnStart'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnStop'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnTick'): bftrader_pb2.BfVoid.FromString,
    ('bftrader.bfrobot.BfRobotService', 'OnTrade'): bftrader_pb2.BfVoid.FromString,
  }
  cardinalities = {
    'OnInit': cardinality.Cardinality.UNARY_UNARY,
    'OnOrder': cardinality.Cardinality.UNARY_UNARY,
    'OnPing': cardinality.Cardinality.UNARY_UNARY,
    'OnStart': cardinality.Cardinality.UNARY_UNARY,
    'OnStop': cardinality.Cardinality.UNARY_UNARY,
    'OnTick': cardinality.Cardinality.UNARY_UNARY,
    'OnTrade': cardinality.Cardinality.UNARY_UNARY,
  }
  stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
  return beta_implementations.dynamic_stub(channel, 'bftrader.bfrobot.BfRobotService', cardinalities, options=stub_options)
# @@protoc_insertion_point(module_scope)
