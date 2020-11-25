/**
 *@Description
 *@ClassName etcd
 *@Date 2020/11/23 5:29 下午
 *@Author ckhero
 */

package plugin

import (
	"base-demo/pkg/config"
	"context"
	"fmt"
	lgo2 "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type Etcd struct {
	registar *etcdv3.Registrar
}

func NewPluginEtcd() *Etcd {
	return &Etcd{
	}
}

func (e *Etcd) InitPlugin() error {
	rc := config.AppConfig.Registry
	options := etcdv3.ClientOptions{
		DialTimeout:  time.Second * time.Duration(rc.DialTimeout),
		DialKeepAlive: time.Second * time.Duration(rc.DialKeepAlive),
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), config.AppConfig.Registry.Address, options)
	if err != nil {
		return err
	}
	Registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%s",config.AppConfig.Registry.Name, rc.Address),
		Value: rc.GrpcAddr,
	}, lgo2.NewNopLogger())
	Registar.Register()
	e.registar = Registar
	return nil
}

func(e *Etcd) Release()  {
	e.registar.Deregister()
	fmt.Println("etcd release")

}