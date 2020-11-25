package handlers

import (
	"context"

	pb "base-demo/truss-demo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GreeterServer {
	return greeterService{}
}

type greeterService struct{}

func (s greeterService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	var resp pb.HelloResponse
	return &resp, nil
}

func (s greeterService) Buy(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	var resp pb.HelloResponse
	return &resp, nil
}
