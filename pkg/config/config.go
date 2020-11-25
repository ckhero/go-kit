/**
 *@Description
 *@ClassName config
 *@Date 2020/11/23 2:16 下午
 *@Author ckhero
 */

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Project  string              `yaml:"project"`
	Application  string             `yaml:"application"`
	Database map[string]Database `yaml:"database"`
	Redis    map[string]Redis               `yaml:"redis"`
	Registry Registry            `yaml:"registry"`
	Domain   map[string]string   `yaml:"domain"`
	Jaeger   Jaeger   		`yaml:"jaeger"`
	Logger   Logger   		`yaml:"logger"`
}

/**
数据库配置
 */
type Database struct {
	Dialect        string
	Database       string
	Username       string
	Password       string
	Host           string
	Port           int
	Charset        string
	MaxIdleConnNum int
	MaxOpenConnNum int
}

type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint16 `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	Database uint8  `json:"database" yaml:"database"`
	MaxIdle int `json:"maxIdle" yaml:"maxIdle"`  // 空闲连接数大小
	MaxActive int `json:"maxActive" yaml:"maxActive"`  // 最大连接数
	IdleTimeout int `json:"idleTimeout" yaml:"idleTimeout"` //空闲连接超时时间
	Prefix string `json:"prefix" yaml:"prefix"` //前缀使用系统名字缩写
}

type Registry struct {
	Name     string
	Address  []string
	DialTimeout      int
	DialKeepAlive int
	GrpcAddr string
}

type Jaeger struct {
	Name     string
	Host  string
	Port      int
}


type Logger struct {
	Level string
}

var (
	AppConfig *Config
	configOnce sync.Once
)

func InitConfig(path string) *Config {
	configOnce.Do(func() {
		AppConfig = loadConfig(path)
	})
	return AppConfig
}

func loadConfig(path string) *Config {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()     // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	ac := &Config{}
	err = viper.Unmarshal(ac)        // 将配置信息绑定到结构体上
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config file %s \n", err))
	}
	return ac
}