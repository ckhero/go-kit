/**
 *@Description
 *@ClassName client
 *@Date 2020/11/22 3:11 下午
 *@Author ckhero
 */

package main

import (
	"base-demo/pkg/config"
	"base-demo/pkg/constant"
	"base-demo/pkg/router/grpc/middleware"
	micro "base-demo/pkg/server/grpc"
	"base-demo/pkg/server/plugin"
	"base-demo/pkg/util"
	"base-demo/protos"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	sdetcd "github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"
	"io"
	"time"
)

func init() {
	path := util.GetArg(constant.ArgConfig, "./config/dev2.yaml")
	config.InitConfig(path)
}

func main() {
	ms := micro.NewRpcServer(
		//plugin.NewPluginRedis(),
		//plugin.NewPluginEtcd(),
		plugin.NewPluginJaeger(),
	)
	defer ms.Release()

	go ms.Run()
	//grpcAddr := "0.0.0.0:5000"
	//ctx := context.WithValue(context.Background(), "ckhero", "ckhero")
	//
	//cc, err := sdetcd.NewClient(ctx, config.AppConfig.Registry.Address, sdetcd.ClientOptions{})
	//entries, err := cc.GetEntries(config.AppConfig.Registry.Name)
	//fmt.Println(entries)
	//conn, err := grpc.Dial(entries[0], grpc.WithInsecure(), grpc.WithUnaryInterceptor(middleware.JaegerClientMiddleware(plugin.GetTracer())))
	//if err != nil {
	//	fmt.Printf("创建grpc连接失败! Error: %s", err)
	//}
	//defer conn.Close()

	Test()

	//Create(conn)
	//	Delete(conn)
}

func Create(conn *grpc.ClientConn) {
	c := protos.NewGreeterClient(conn)
	rsp, err := c.Hello(context.TODO(),	&protos.HelloRequest{Name:"11d1"})
	fmt.Println(err, rsp)
	//
	//c1 := protos.NewDemoClient(conn)
	//rsp1, err := c1.Hello1(context.TODO(),	&protos.HelloRequest2{Name:"11d1"})
	//fmt.Println(err, rsp1)
	//rsp2, err := c1.Hello2(context.TODO(),	&protos.HelloRequest1{Name:"11d1"})
	//fmt.Println(err, rsp2)
	//in := &protos.HelloRequest{
	//	Name: "wss",
	//}
	//greeterSrv := service.NewGreeterSvc()
	//endpoints := endpoint.NewEndpoints(greeterSrv)
	//grpc2.NewGreeterClient(endpoints, conn)
	//out := &protos.HelloResponse{}
	//ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	//err := conn.Invoke(ctx, "/protos.Greeter/Hello", in, out)
	//
	//fmt.Println(err)
	//fmt.Println(out)
}

func Test() {
	ctx := context.Background()
	//连接注册中心
	client, err := sdetcd.NewClient(ctx, config.AppConfig.Registry.Address, sdetcd.ClientOptions{})

	if err != nil {
		panic(err)
	}
	logger := log.NewNopLogger()
	//创建实例管理器, 此管理器会Watch监听etc中prefix的目录变化更新缓存的服务实例数据
	instancer, err := sdetcd.NewInstancer(client, config.AppConfig.Registry.Name, logger)
	if err != nil {
		panic(err)
	}
	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, reqFactory, logger) //reqFactory自定义的函数，主要用于端点层（endpoint）接受并显示数据
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)

	/**
	我们可以通过负载均衡器直接获取请求的endPoint，发起请求
	reqEndPoint,_ := balancer.Endpoint()
	*/

	/**
	也可以通过retry定义尝试次数进行请求
	*/
	ch := make(chan struct{})
	ch <- struct{}{}
	reqEndPoint := lb.Retry(3, 10*time.Second, balancer)

	//现在我们可以通过 endPoint 发起请求了
	req := struct{}{}
	ctx = context.WithValue(ctx, "test", "test")
	if _, err = reqEndPoint(ctx, req); err != nil {
		panic(err)
	}

}


func reqFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("请求服务: ", instanceAddr)
		conn, err := grpc.Dial(instanceAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(middleware.JaegerClientMiddleware(plugin.GetTracer())))
		if err != nil {
			fmt.Println(err)
			panic("connect error")
		}
		defer conn.Close()
		bookClient := protos.NewGreeterClient(conn)
		bi, _ := bookClient.Hello(ctx, &protos.HelloRequest{
			Name:                 "name",
		})
		fmt.Println("获取书籍详情")
		fmt.Println(bi)

		return nil, nil
	}, nil, nil
}