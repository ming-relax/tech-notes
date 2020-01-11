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
	for {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}

		c := pb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err = c.HelloUnary(ctx, &pb.HelloRequest{Name: "hello"})
		if err != nil {
			conn.Close()
			log.Fatalf("failed to call: %v", err)
		}
		cancel()
		conn.Close()
	}
}