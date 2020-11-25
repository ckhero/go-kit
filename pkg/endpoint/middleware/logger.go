/**
 *@Description
 *@ClassName log
 *@Date 2020/11/25 10:50 上午
 *@Author ckhero
 */

package middleware

import (
	"base-demo/pkg/log/logger"
	"context"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func MakeLoggerEndpointMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.GetLogger(ctx).Info(time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}