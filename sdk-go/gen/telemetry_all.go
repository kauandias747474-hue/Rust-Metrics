package telemetry

import (
	"context"

	"google.golang.org/grpc"
)

// --- DEFINIÇÕES DAS MENSAGENS ---

type Empty struct{}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return "empty" }
func (*Empty) ProtoMessage()    {}

type MetricRequest struct {
	SensorId  string  `protobuf:"bytes,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	Value     float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MetricRequest) Reset()         { *m = MetricRequest{} }
func (m *MetricRequest) String() string { return "metric_request" }
func (*MetricRequest) ProtoMessage()    {}

type MetricResponse struct {
	Success           bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	CalculatedStdDev float32 `protobuf:"fixed32,2,opt,name=calculated_std_dev,json=calculatedStdDev,proto3" json:"calculated_std_dev,omitempty"`
	Message           string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MetricResponse) Reset()         { *m = MetricResponse{} }
func (m *MetricResponse) String() string { return "metric_response" }
func (*MetricResponse) ProtoMessage()    {}

type StatusResponse struct {
	Message        string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CpuUsage       string `protobuf:"bytes,2,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	HardwareLinked bool   `protobuf:"varint,3,opt,name=hardware_linked,json=hardwareLinked,proto3" json:"hardware_linked,omitempty"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return "status_response" }
func (*StatusResponse) ProtoMessage()    {}

type BatchRequest struct {
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return "batch_request" }
func (*BatchRequest) ProtoMessage()    {}

type BatchResponse struct {
	Metrics []*MetricRequest `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (m *BatchResponse) Reset()         { *m = BatchResponse{} }
func (m *BatchResponse) String() string { return "batch_response" }
func (*BatchResponse) ProtoMessage()    {}

// --- LÓGICA DO CLIENTE gRPC ---

type MetricsIngestorClient interface {
	SendMetricsStream(ctx context.Context, opts ...grpc.CallOption) (MetricsIngestor_SendMetricsStreamClient, error)
	GetSystemStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type metricsIngestorClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsIngestorClient(cc grpc.ClientConnInterface) MetricsIngestorClient {
	return &metricsIngestorClient{cc}
}

func (c *metricsIngestorClient) SendMetricsStream(ctx context.Context, opts ...grpc.CallOption) (MetricsIngestor_SendMetricsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &grpc.StreamDesc{
		StreamName:    "SendMetricsStream",
		ClientStreams: true,
	}, "/telemetry.MetricsIngestor/SendMetricsStream", opts...)
	if err != nil {
		return nil, err
	}
	return &metricsIngestorSendMetricsStreamClient{stream}, nil
}

type MetricsIngestor_SendMetricsStreamClient interface {
	Send(*MetricRequest) error
	CloseAndRecv() (*MetricResponse, error)
	grpc.ClientStream
}

type metricsIngestorSendMetricsStreamClient struct {
	grpc.ClientStream
}

func (x *metricsIngestorSendMetricsStreamClient) Send(m *MetricRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsIngestorSendMetricsStreamClient) CloseAndRecv() (*MetricResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MetricResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsIngestorClient) GetSystemStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/telemetry.MetricsIngestor/GetSystemStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
