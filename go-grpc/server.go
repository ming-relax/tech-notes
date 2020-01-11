package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"net/http"
	"perf-grpc/pb"
)

const (
	port = ":5051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) HelloUnary(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.GetName()}, nil
}

func (s *server) HelloBiStream(srv pb.Greeter_HelloBiStreamServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		// simulate unary rpc, start a go routine to handle this.
		go func(reqParam *pb.HelloRequest) {
			err = srv.Send(&pb.HelloReply{Message: reqParam.GetName()})
			if err != nil {
				log.Fatalf("send error: %v", err)
			}
		}(req)
	}
}

func main() {
	ls, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}


	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	pb.RegisterGreeterServer(s, &server{})

	grpc_prometheus.Register(s)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			log.Fatalf("failed to listen and serve /metrics")
		}
	}()

	if err := s.Serve(ls); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}