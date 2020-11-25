/**
 *@Description
 *@ClassName jaeger
 *@Date 2020/11/24 11:24 下午
 *@Author ckhero
 */

package middleware

import (
	"github.com/opentracing/opentracing-go"
)

type tracerMiddlewareServer struct {
	next   interface{}
	tracer opentracing.Tracer
}
type NewMiddlewareSvc func(svc interface{}) interface{}

func NewTracerMiddlewareSvc(tracer opentracing.Tracer) NewMiddlewareSvc {
	return func(service interface{}) interface{} {
		return tracerMiddlewareServer{
			next:   service,
			tracer: tracer,
		}
	}
}

//func (l tracerMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
//	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, l.tracer, "service", opentracing.Tag{
//		Key:   string(ext.Component),
//		Value: "NewTracerServerMiddleware",
//	})
//	defer func() {
//		span.Finish()
//	}()
//	out, err = l.next.Login(ctxContext, in)
//	return
//}
