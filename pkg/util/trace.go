/**
 *@Description
 *@ClassName trace
 *@Date 2020/11/25 10:49 上午
 *@Author ckhero
 */

package util

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

/**
 * 获取链路tranceId
 */
func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if sp, ok := span.Context().(jaeger.SpanContext); ok {
			return sp.TraceID().String()
		}
		return ""
	}

	return ""
}

/**
 * 设置span标签
 */
func SetSpanTag(ctx context.Context, key string, value interface{}) {
	if ctx == nil {
		return
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag(key, value)
	}
}