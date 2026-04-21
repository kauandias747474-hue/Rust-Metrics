package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"sdk-go/gen"   
	"sdk-go/metrics"
)

type RustClient struct {
	Conn   *grpc.ClientConn         
	Client gen.MetricsIngestorClient 
}

func InitClient(target string) (*RustClient, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &RustClient{
		Conn:   conn,
		Client: gen.NewMetricsIngestorClient(conn),
	}, nil
}

func (c *RustClient) EnviarParaMotor(m metrics.SystemMetrics) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.Client.SendMetricsStream(ctx)
	if err != nil {
		log.Printf("Erro ao abrir stream: %v", err)
		return
	}

	req := &gen.MetricRequest{
		SensorId:  m.SensorID,
		Value:     float32(m.Value),
		Timestamp: time.Now().Unix(),
	}

	stream.Send(req)
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Erro na resposta: %v", err)
		return
	}
	log.Printf(" Motor Rust diz: %s", res.Message)
}
