/**
 *@Description
 *@ClassName endpoint
 *@Date 2020/11/22 3:24 下午
 *@Author ckhero
 */

package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	// greeter
	GreeterHelloEndpoint endpoint.Endpoint
	//demo
	DemoHello1Enpoint endpoint.Endpoint
	DemoHello2Enpoint endpoint.Endpoint
}

func WrapEndpoints(in interface{}) interface{} {

	return in
}

func WrapService(in interface{}) interface{} {
	return in
}