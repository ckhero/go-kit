// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: ef2331b7e2
// Version Date: 2020-10-07T23:22:38Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "base-demo/truss-demo"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC GreeterServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.GreeterServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// greeter

		hello: grpctransport.NewServer(
			endpoints.HelloEndpoint,
			DecodeGRPCHelloRequest,
			EncodeGRPCHelloResponse,
			serverOptions...,
		),
		buy: grpctransport.NewServer(
			endpoints.BuyEndpoint,
			DecodeGRPCBuyRequest,
			EncodeGRPCBuyResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the GreeterServer interface
type grpcServer struct {
	hello grpctransport.Handler
	buy   grpctransport.Handler
}

// Methods for grpcServer to implement GreeterServer interface

func (s *grpcServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, rep, err := s.hello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResponse), nil
}

func (s *grpcServer) Buy(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, rep, err := s.buy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResponse), nil
}

// Server Decode

// DecodeGRPCHelloRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC hello request to a user-domain hello request. Primarily useful in a server.
func DecodeGRPCHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	return req, nil
}

// DecodeGRPCBuyRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC buy request to a user-domain buy request. Primarily useful in a server.
func DecodeGRPCBuyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCHelloResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain hello response to a gRPC hello reply. Primarily useful in a server.
func EncodeGRPCHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HelloResponse)
	return resp, nil
}

// EncodeGRPCBuyResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain buy response to a gRPC buy reply. Primarily useful in a server.
func EncodeGRPCBuyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HelloResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}