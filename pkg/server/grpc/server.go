/**
 *@Description
 *@ClassName service
 *@Date 2020/11/23 4:42 下午
 *@Author ckhero
 */

package grpc

import (
	"base-demo/pkg/config"
	"base-demo/pkg/log/logger"
	router "base-demo/pkg/router/grpc"
	"base-demo/pkg/server/plugin"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type RpcServer struct {
	plugins []plugin.Plugin
	Errc chan error
}


func NewRpcServer(plugins ...plugin.Plugin) *RpcServer {
	defaultPlugins := []plugin.Plugin{		plugin.NewPluginLogger()}
	plugins = append(defaultPlugins, plugins...)
	return &RpcServer{
		plugins: plugins,
		Errc:    make(chan error),
	}
}

func (ms *RpcServer) Run()  {
	// 初始化插件列表
	plugin.InitPlugin(ms.plugins...)
	// 初始化grpc链接
	ln, err := initGrpcListen()
	if err != nil {
		panic(fmt.Errorf("Fatal init grpc: %s \n", err))
	}
	// 注册服务
	router.Register(ln, ms.Errc)
	log.WithField("error", <- ms.Errc).Info("Exit")
}

func (ms *RpcServer) Release()  {
	for _, p := range ms.plugins {
		p.Release()
	}
}

func initGrpcListen() (net.Listener, error){
	grpcAddr := config.AppConfig.Registry.GrpcAddr
	logger.GetLogger(context.Background()).WithField("address", grpcAddr).Info("grpc start")
	ln, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, err
	}
	return ln, nil
}



