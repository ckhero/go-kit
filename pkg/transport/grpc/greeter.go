/**
 *@Description
 *@ClassName greeter
 *@Date 2020/11/22 2:19 下午
 *@Author ckhero
 */

package grpc

import (
	"base-demo/pkg/endpoint"
	"base-demo/protos"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type greeterServer struct {
	hello grpctransport.Handler
}

func (u greeterServer) Hello(c context.Context, req *protos.HelloRequest) (*protos.HelloResponse, error) {
	_, rep, err := u.hello.ServeGRPC(c, req)
	if err != nil {
		return nil, err
	}
	return rep.(*protos.HelloResponse), nil
}

func NewGreeterServer(endpoints endpoint.GreeterEndpoints, opts ...grpctransport.ServerOption) interface{} {
	return greeterServer{
		hello:grpctransport.NewServer(
				endpoints.GreeterHelloEndpoint,
				DecodeRequest,
				EncodeResponse,
				opts...
			),
	}
}

// Server
// 1. decode request          pb -> model
func DecodeRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

// 2. encode response           model -> pb
func EncodeResponse(c context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

