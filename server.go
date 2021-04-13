package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/kimilton/grpcgo/proto"
	"github.com/kimilton/grpcgo/shared"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCommunicatorServer
}

func (s *server) InititateCommunication(ctx context.Context, in *pb.InitRequest) (*pb.Ack, error) {
	ack := &pb.Ack{Ack: true}
	if in.Name == shared.DEFAULT_NAME {
		ack.Ack = false
	}
	fmt.Printf("[Request] Received from %s | Ack: %v\n", in.Name, ack.Ack)
	return ack, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunicatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
