package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"perf-grpc/pb"
	//"time"
)

const (
	address = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	stream, err := c.HelloBiStream(context.Background())
	if err != nil {
		log.Fatalf("failed to create stream: %v", err)
	}
	//defer cancel()

	for {
		if err := stream.Send(&pb.HelloRequest{Name: "hello"}); err != nil {
			log.Fatalf("failed to send msg: %v", err)
		}

		if _, err := stream.Recv(); err != nil {
			log.Fatalf("failed to receive msg: %v", err)
		}

	}
}