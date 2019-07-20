package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// GetMeasurements implements MeterUsageServer
func (s *server) GetMeasurements(request *GetMeasurementsRequest, stream MeterUsage_GetMeasurementsServer) error {
	measurement := Measurement{
		Timestamp:  1234,
		Metervalue: 5678,
	}
	if err := stream.Send(&measurement); err != nil {
		return err
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterMeterUsageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
