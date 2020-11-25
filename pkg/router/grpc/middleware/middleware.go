/**
 *@Description
 *@ClassName middleware
 *@Date 2020/11/22 2:52 下午
 *@Author ckhero
 */

package middleware

import (
	"base-demo/pkg/log/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func RecoveryInterceptorMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
			fmt.Println(err)
		}
	}()

	return handler(ctx, req)
}

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.GetLogger(ctx).WithField("req", req).WithField("method", info).Info("request ! ")
	resp, err = handler(ctx, req)
	logger.GetLoggerWithBody(ctx, resp).Info("response")
	return
}

