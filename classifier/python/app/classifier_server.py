from concurrent import futures
import logging
import sys
import grpc

import classify_pb2
import classify_pb2_grpc
import classifier

class ClassifierServicer(classify_pb2_grpc.ClassifierServicer):
    def Classify(self, request, context):
        tf = classifier.ToiletFinder(predict=True)
        return classify_pb2.Result(name=str(tf.predict(request.path)))

def serve(port: str):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    classify_pb2_grpc.add_ClassifierServicer_to_server(
        ClassifierServicer(), server)
    server.add_insecure_port(port)
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    port = sys.argv[1]
    if port == '':
        port = '[::]:7771'
    logging.basicConfig()
    serve(port)
