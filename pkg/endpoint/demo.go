/**
 *@Description
 *@ClassName demo
 *@Date 2020/11/24 2:20 下午
 *@Author ckhero
 */

package endpoint

import (
	"base-demo/pkg/endpoint/middleware"
	"base-demo/pkg/errors"
	"base-demo/protos"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type DemoEndpoints struct {
	DemoHello1Endpoint endpoint.Endpoint
	DemoHello2Endpoint endpoint.Endpoint
}
func NewDemoEndpoints(service protos.DemoServer) DemoEndpoints {
	// Business domain.
	endpoints := DemoEndpoints{
		DemoHello1Endpoint: MakeDemoHello1Endpoint(service),
		DemoHello2Endpoint: MakeDemoHello2Endpoint(service),
	}

	// Wrap selected Endpoints with middlewares. See handlers/middlewares.go
	endpoints = demoWrapEndpoints(endpoints)

	return endpoints
}

func demoWrapEndpoints(in DemoEndpoints) DemoEndpoints {
	in.DemoHello1Endpoint = middleware.MakeTracerEndpointMiddleware()(in.DemoHello1Endpoint)
	in.DemoHello1Endpoint = middleware.MakeTracerEndpointMiddleware()(in.DemoHello1Endpoint)
	return in
}

func MakeDemoHello1Endpoint(svc protos.DemoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*protos.HelloRequest2)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Hello1(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeDemoHello2Endpoint(svc protos.DemoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*protos.HelloRequest1)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Hello2(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}