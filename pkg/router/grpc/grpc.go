/**
 *@Description
 *@ClassName grcp
 *@Date 2020/11/24 2:02 下午
 *@Author ckhero
 */

package grpc

import (
	"base-demo/pkg/router/grpc/middleware"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

var defaultMiddlewareOpts = []grpc.ServerOption{

	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		middleware.JaegerServerMiddleware(),
		grpctransport.Interceptor,
		middleware.LoggerMiddleware,
		middleware.RecoveryInterceptorMiddleware,
	)),
}

func Register(ln net.Listener, errc chan error) {
	// 初始化grpc服务
	grpcSrv := grpc.NewServer(defaultMiddlewareOpts...)

	// 注册 欢迎服务
	RegisterGreeterService(grpcSrv)
	// 注册 demo
	RegisterDemoService(grpcSrv)

	// 开始服务
	errc <- grpcSrv.Serve(ln)
}