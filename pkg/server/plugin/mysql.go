/**
 *@Description
 *@ClassName mysql
 *@Date 2020/11/25 9:25 上午
 *@Author ckhero
 */

package plugin

import (
	"base-demo/pkg/config"
	"base-demo/pkg/db/mysql"
	"fmt"
)

type Mysql struct {
}

func NewPluginMysql() *Mysql {
	return &Mysql{
	}
}

func (r *Mysql) InitPlugin() error {
	mysql.ConnectDB(config.AppConfig.Database)
	return nil
}

func (r *Mysql) Release() {
	mysql.CloseDB()
	fmt.Println("redis release")

}