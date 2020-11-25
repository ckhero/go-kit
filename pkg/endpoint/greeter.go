/**
 *@Description
 *@ClassName greeter
 *@Date 2020/11/22 2:14 下午
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

type GreeterEndpoints struct {
	GreeterHelloEndpoint endpoint.Endpoint
}

func NewGreeterEndpoints(svc protos.GreeterServer) GreeterEndpoints {
	eds := GreeterEndpoints{
		GreeterHelloEndpoint: MakeGreeterHelloEndpoint(svc),
	}
	eds = wrapEndpoints(eds)
	return eds
}

func wrapEndpoints(in GreeterEndpoints) GreeterEndpoints {
	in.GreeterHelloEndpoint = middleware.MakeTracerEndpointMiddleware()(in.GreeterHelloEndpoint)
	in.GreeterHelloEndpoint = middleware.MakeLoggerEndpointMiddleware()(in.GreeterHelloEndpoint)
	return in
}

func MakeGreeterHelloEndpoint(svc protos.GreeterServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(*protos.HelloRequest)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Hello(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
