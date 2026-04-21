import grpc
import sys
import os

sys.path.append(os.path.dirname(os.path.abspath(__file__)))

import telemetry_pb2
import telemetry_pb2_grpc

def run():

    channel = grpc.insecure_channel('localhost:50051')
  
    stub = telemetry_pb2_grpc.MetricsIngestorStub(channel)

    try:
      
        request = telemetry_pb2.MetricRequest(sensor_id="PYTHON_01", value=25.5)
      
        print("Tentando enviar dado para o Rust...")
        response = stub.IngestMetric(request, timeout=5) 
        
        print(f" Resposta: {response.message}")

    except grpc.RpcError as e:
        print(f" Erro de RPC: {e.code()}")
        print("DICA: O servidor Rust (localhost:50051) está ligado?")
    except Exception as e:
        print(f" Outro erro: {e}")

if __name__ == '__main__':
    run()
