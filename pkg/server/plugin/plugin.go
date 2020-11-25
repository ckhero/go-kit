/**
 *@Description
 *@ClassName plugin
 *@Date 2020/11/24 1:38 下午
 *@Author ckhero
 */

package plugin

import "fmt"

type Plugin interface {
	InitPlugin() error
	Release()
}

func InitPlugin(plugins ...Plugin) {
	for _, p := range plugins {
		err := p.InitPlugin()
		if err != nil {
			panic(fmt.Errorf("Fatal init server: %s \n", err))
		}
	}
}
