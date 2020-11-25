/**
 *@Description
 *@ClassName logger
 *@Date 2020/11/25 12:52 上午
 *@Author ckhero
 */

package middleware

import log "github.com/sirupsen/logrus"

type Service struct {

}
type ServiceMiddleware func(Service) Service


type loggingMiddleware struct {
	Service
	logger log.Logger
}
//
//func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
//	return func(next Service) Service {
//		return loggingMiddleware{next, logger}
//	}
//}