/**
 *@Description
 *@ClassName logger
 *@Date 2020/11/25 10:58 上午
 *@Author ckhero
 */

package plugin

import (
	"base-demo/pkg/config"
	"base-demo/pkg/log/logger"
)

type Logger struct {
}

func NewPluginLogger() *Logger {
	return &Logger{
	}
}

func (r *Logger) InitPlugin() error {
	appConfig := config.AppConfig
	logger.InitLogger(appConfig.Project, appConfig.Application, appConfig.Logger)
	return nil
}

func (r *Logger) Release() {

}