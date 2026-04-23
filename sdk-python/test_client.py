import grpc
import telemetry_pb2
import telemetry_pb2_grpc
import time

def run():
   
    with grpc.insecure_channel('127.0.0.1:50051') as channel:
        stub = telemetry_pb2_grpc.MetricsIngestorStub(channel)
        
 
    
        dados = [10.0, 20.0, 30.0]
        
        def generate_metrics():
            for val in dados:
                yield telemetry_pb2.MetricRequest(
                    sensor_id="SENSOR_TESTE",
                    value=val,
                    timestamp=int(time.time())
                )

        response = stub.SendMetricsStream(generate_metrics())
        
        print(f" Resposta do Servidor: {response.message}")
        print(f" Desvio Padrão (C++): {response.calculated_std_dev}")

if __name__ == '__main__':
    run()
