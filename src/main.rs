mod common;
mod ingestion;
mod storage;

use std::thread;
pub mod telemetry {
    tonic::include_proto!("telemetry");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
  
   
    thread::spawn(|| {
        ingestion::udp::start_udp_server();
    });

    println!(" Servidor gRPC rodando em [::1]:50051");
    println!(" Servidor UDP escutando em 127.0.0.1:8080");

    let addr = "[::1]:50051".parse()?;
    let sink = storage::DataSink::new(); // Isso faz o aviso do DataSink sumir!
    let metrics_service = ingestion::grpc::MyMetrics { sink };

   
    tonic::transport::Server::builder()
        .add_service(telemetry::metrics_ingestor_server::MetricsIngestorServer::new(metrics_service))
        .serve(addr)
        .await?;

    Ok(())
}
