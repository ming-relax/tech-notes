package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"perf-grpc/pb"
	"time"
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
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err = c.HelloUnary(ctx, &pb.HelloRequest{Name: "hello"})
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		cancel()
	}
}