/**
 *@Description
 *@ClassName demo
 *@Date 2020/11/24 2:24 下午
 *@Author ckhero
 */

package grpc

import (
	"base-demo/pkg/endpoint"
	"base-demo/protos"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type demoServer struct {
	hello1 grpctransport.Handler
	hello2 grpctransport.Handler
}

func (u demoServer) Hello1(c context.Context, req *protos.HelloRequest2) (*protos.HelloResponse2, error) {
	_, rep, err := u.hello1.ServeGRPC(c, req)
	if err != nil {
		return nil, err
	}
	return rep.(*protos.HelloResponse2), nil
}

func (u demoServer) Hello2(c context.Context, req *protos.HelloRequest1) (*protos.HelloResponse1, error) {
	_, rep, err := u.hello2.ServeGRPC(c, req)
	if err != nil {
		return nil, err
	}
	return rep.(*protos.HelloResponse1), nil
}

func NewDemoServer(eds endpoint.DemoEndpoints, opts ...grpctransport.ServerOption) protos.DemoServer {
	return demoServer{
		hello2:grpctransport.NewServer(
			eds.DemoHello1Endpoint,
			DecodeRequest,
			EncodeResponse,
			opts...
		),
		hello1:grpctransport.NewServer(
			eds.DemoHello2Endpoint,
			DecodeRequest,
			EncodeResponse,
			opts...
		),

	}
}