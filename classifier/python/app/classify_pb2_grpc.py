# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import classify_pb2 as classify__pb2


class ClassifierStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Classify = channel.unary_unary(
        '/model.Classifier/Classify',
        request_serializer=classify__pb2.Image.SerializeToString,
        response_deserializer=classify__pb2.Result.FromString,
        )


class ClassifierServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Classify(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ClassifierServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Classify': grpc.unary_unary_rpc_method_handler(
          servicer.Classify,
          request_deserializer=classify__pb2.Image.FromString,
          response_serializer=classify__pb2.Result.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'model.Classifier', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
