package main

import (
	"context"
	"log"
	"net"
	pb "root/proto"

	"google.golang.org/grpc"
)

type test struct {
	pb.UnimplementedGetHelloServer
}

func (s *test) HelloWorld(ctx context.Context, in *pb.HelloWorldResponse) (*pb.HelloWorldResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.HelloWorldResponse{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetHelloServer(s, &test{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
