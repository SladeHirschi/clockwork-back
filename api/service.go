package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "clockwork/protos/clockwork/v1" // Import your generated gRPC package
)

type ClockworkApi struct {
	pb.UnimplementedClockWorkServiceServer
	vipe *viper.Viper
}

func NewClockworkApi(vipe *viper.Viper) *ClockworkApi {
	return &ClockworkApi{
		vipe: vipe,
	}
}

func (c *ClockworkApi) Serve() {
	port := c.vipe.GetString("server.port") // Get port from config
	if port == "" {
		port = "50051" // Default gRPC port
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register your gRPC service
	pb.RegisterClockWorkServiceServer(grpcServer, c) // Ensure ClockworkApi implements the service interface

	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (c *ClockworkApi) ClockIn(ctx context.Context, req *pb.ClockInRequest) (*pb.ClockInResponse, error) {
	fmt.Println("Something")
	return nil, nil
}
