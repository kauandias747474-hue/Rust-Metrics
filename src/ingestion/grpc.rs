use tonic::{Request, Response, Status};
use crate::telemetry::metrics_ingestor_server::MetricsIngestor;
use crate::telemetry::{
    MetricRequest, MetricResponse, Empty, 
    StatusResponse, BatchRequest, BatchResponse
};
use crate::storage::DataSink;

pub struct MyMetrics {
    pub sink: DataSink,
}

#[tonic::async_trait]
impl MetricsIngestor for MyMetrics {
 
    async fn send_metrics_stream(
        &self,
         request: Request<tonic::Streaming<MetricRequest>>,
    ) -> Result<Response<MetricResponse>, Status> {
        let mut stream = request.into_inner();

        while let Ok(Some(metric)) = stream.message().await {
            // Garanta que save_metric existe no DataSink
            self.sink.save_metric(&metric);
        }

        Ok(Response::new(MetricResponse {
            success: true,
            calculated_std_dev: 0.0, 
            message: "Métricas de stream processadas".into(),
        }))
    }


    async fn get_system_status(
        &self, 
        _request: Request<Empty>
    ) -> Result<Response<StatusResponse>, Status> {
        Ok(Response::new(StatusResponse {
            message: "Motor Rust Ativo".into(),
            cpu_usage: "10%".into(),
            hardware_linked: true,
        }))
    }

    // Função obrigatória do contrato
    async fn get_metrics_batch(
        &self, 
        _request: Request<BatchRequest>
    ) -> Result<Response<BatchResponse>, Status> {
        Ok(Response::new(BatchResponse {
            metrics: vec![],
        }))
    }
}
