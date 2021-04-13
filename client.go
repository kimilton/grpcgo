package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/kimilton/grpcgo/proto"
	"github.com/kimilton/grpcgo/shared"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	name := shared.DEFAULT_NAME
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	c := pb.NewCommunicatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.InitRequest{Name: name}

	ack, err := c.InititateCommunication(ctx, req)
	if err != nil {
		log.Fatalf("Could not initiate communication: %v", err)
	}

	if ack.Ack == true {
		fmt.Println("Acknowledged")
	} else {
		fmt.Println("Rejected")
	}

}
