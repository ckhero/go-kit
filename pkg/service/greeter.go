/**
 *@Description
 *@ClassName greeter
 *@Date 2020/11/22 2:09 下午
 *@Author ckhero
 */

package service

import (
	"base-demo/pkg/dal/dao"
	"base-demo/pkg/server/plugin"
	"base-demo/protos"
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type greeterSvc struct {

}

func NewGreeterSvc() protos.GreeterServer {
	svc := &greeterSvc{}
	return svc
}

func(g *greeterSvc) Hello(ctx context.Context, req *protos.HelloRequest) (*protos.HelloResponse, error) {
	span, _ := opentracing.StartSpanFromContextWithTracer(ctx, plugin.GetTracer(), "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer span.Finish()
	_ = dao.NewAreaDao().Create(ctx)
	//fmt.Println(res)
	//fmt.Println(err)
	return &protos.HelloResponse{
		Code: 0,
		Msg:  "",
		Data: &protos.HelloResponseData{Greeting:fmt.Sprintf("fsdafadsfas %s", req.Name)},
	}, errors.New("fasdfasd")
}

func(g *greeterSvc) Buy(ctx context.Context, req *protos.HelloRequest) (*protos.HelloResponse, error) {
	return nil, nil
}