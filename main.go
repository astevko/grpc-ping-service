package grpcping

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/astevko/grpc-ping-service/pkg/grpcping"

	"google.golang.org/grpc"
)

// server implements the PingService server interface
type server struct{}

// Ping implements the PingService Ping method
func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "OK",
	}, nil
}

func main() {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register our server implementation with the gRPC server
	pb.RegisterPingServiceServer(grpcServer, &server{})

	// Start the gRPC server on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("gRPC server listening on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
