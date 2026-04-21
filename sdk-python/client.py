
import grpc
import telemetry_pb2
import telemetry_pb2_grpc

def run():
 
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = telemetry_pb2_grpc.MetricsIngestorStub(channel)
        
     
        request = telemetry_pb2.MetricRequest(
            sensor_id="PYTHON_SENSOR_01", 
            value=24.5
        )
        
      
        response = stub.IngestMetric(request)
        print(f"Resposta do Motor Rust: {response.message}")

if __name__ == '__main__':
    run()
