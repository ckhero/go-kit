/**
 *@Description
 *@ClassName redis
 *@Date 2020/11/23 4:59 下午
 *@Author ckhero
 */

package plugin

import (
	"base-demo/pkg/config"
	"base-demo/pkg/db/redis"
	"fmt"
)

type Redis struct {
}

func NewPluginRedis() *Redis {
	return &Redis{
	}
}

func (r *Redis) InitPlugin() error {
	for k, v := range config.AppConfig.Redis {
		redis.ConnectRedis(k, v)
	}
	return nil
}

func (r *Redis) Release() {
	redis.CloseRedis()
	fmt.Println("redis release")

}