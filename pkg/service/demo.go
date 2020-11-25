/**
 *@Description
 *@ClassName demo
 *@Date 2020/11/24 2:13 下午
 *@Author ckhero
 */

package service

import (
	"base-demo/pkg/errors"
	"base-demo/protos"
	"context"
	"fmt"
)

type demoSvc struct {

}

func NewDemoSvc() protos.DemoServer {
	svc := &demoSvc{}
	return svc
}

func(g *demoSvc) Hello1(ctx context.Context, req *protos.HelloRequest2) (*protos.HelloResponse2, error) {
	return nil, errors.EndpointTypeError
}

func(g *demoSvc) Hello2(ctx context.Context, req *protos.HelloRequest1) (*protos.HelloResponse1, error) {
	return &protos.HelloResponse1{
		Code: 0,
		Msg:  "",
		Data: &protos.HelloResponseData1{Greeting:fmt.Sprintf("fsdafadsfas2222222 %s", req.Name)},
	}, nil
}