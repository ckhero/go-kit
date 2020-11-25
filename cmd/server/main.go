/**
 *@Description
 *@ClassName main
 *@Date 2020/11/22 3:06 下午
 *@Author ckhero
 */

package main

import (
	"base-demo/pkg/config"
	"base-demo/pkg/constant"
	_ "base-demo/pkg/router/grpc"
	micro "base-demo/pkg/server/grpc"
	"base-demo/pkg/server/plugin"
	. "base-demo/pkg/util"
)

func init() {
	path := GetArg(constant.ArgConfig, "./config/dev.yaml")
	config.InitConfig(path)
}

func main() {
	ms := micro.NewRpcServer(
		plugin.NewPluginRedis(),
		plugin.NewPluginEtcd(),
		plugin.NewPluginJaeger(),
		plugin.NewPluginMysql(),
	)
	ms.Run()
	defer ms.Release()
}