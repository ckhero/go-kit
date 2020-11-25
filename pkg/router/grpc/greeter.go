/**
 *@Description
 *@ClassName greeter
 *@Date 2020/11/22 2:42 下午
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



func RegisterGreeterService(grpcServer *grpc.Server) {

	// service 逻辑处理层
	svc := service.NewGreeterSvc()

	// endpoint 层
	endpoints := endpoint.NewGreeterEndpoints(svc)
	// transport 层
	srv := transport.NewGreeterServer(endpoints).(protos.GreeterServer)

	//注册服务
	protos.RegisterGreeterServer(grpcServer, srv)
}

