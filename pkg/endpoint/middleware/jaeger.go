/**
 *@Description
 *@ClassName jaeger
 *@Date 2020/11/24 11:11 下午
 *@Author ckhero
 */

package middleware

import (
	"base-demo/pkg/server/plugin"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func MakeTracerEndpointMiddleware() endpoint.Middleware {
	tracer := plugin.GetTracer()
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "endpoint", opentracing.Tag{
				Key:   string(ext.Component),
				Value: "NewTracerEndpointMiddleware",
			})
			defer span.Finish()
			return next(ctxContext, request)
		}
	}
}