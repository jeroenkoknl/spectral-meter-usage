package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// GetMeasurements implements MeterUsageServer
func (s *server) GetMeasurements(request *GetMeasurementsRequest, stream MeterUsage_GetMeasurementsServer) error {
	log.Print("GetMeasurements")

	// Random meter values for every minute the last 24 hours
	metervalue := rand.Int31()
	for minutes := -24 * 60; minutes < 0; minutes++ {
		timestamp := time.Now().Add(time.Minute * time.Duration(minutes))
		metervalue += rand.Int31n(100)
		measurement := Measurement{
			Timestamp:  timestamp.Unix(),
			Metervalue: metervalue,
		}
		if err := stream.Send(&measurement); err != nil {
			return err
		}
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
