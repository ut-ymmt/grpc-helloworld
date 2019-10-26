package main

import (
	"context"
	pb "github.com/ut-ymmt/grpc-helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
	}
}
