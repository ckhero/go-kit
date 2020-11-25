/**
 *@Description
 *@ClassName demo
 *@Date 2020/11/24 2:19 下午
 *@Author ckhero
 */

package grpc

import (
	"base-demo/pkg/endpoint"
	"base-demo/pkg/service"
	transport "base-demo/pkg/transport/grpc"
	"base-demo/protos"
	"google.golang.org/grpc"
)

func RegisterDemoService(grpcServer *grpc.Server) {

	// service 逻辑处理层
	svc := service.NewDemoSvc()

	// endpoint 层
	endpoints := endpoint.NewDemoEndpoints(svc)

	// transport 层
	srv := transport.NewDemoServer(endpoints)

	//注册服务
	protos.RegisterDemoServer(grpcServer, srv)
}