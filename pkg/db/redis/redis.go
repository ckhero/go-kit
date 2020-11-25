/**
 *@Description
 *@ClassName redis
 *@Date 2020/11/2 4:56 下午
 *@Author ckhero
 */

package redis

import (
	"base-demo/pkg/config"
	"base-demo/pkg/constant"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type rds struct {
	connect redis.Conn
}

var rdsPool *redis.Pool

var rdsPoolMap = make(map[string]*redis.Pool )

var prefix string

/**
 * 连接redis
 */
func ConnectRedis(key string, redisConfig config.Redis) {
	address := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	rdsPool = &redis.Pool{
		Wait:        true,
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		IdleTimeout: time.Duration(redisConfig.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address, redis.DialPassword(redisConfig.Password))
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("SELECT", redisConfig.Database)
			return conn, nil
		},
	}
	rdsPoolMap[key] = rdsPool
	prefix = redisConfig.Prefix
}

/**
 * 关闭redis连接池
 */
func CloseRedis() {
	if rdsPool != nil {
		_ = rdsPool.Close()
	}
}

/**
 * 获取一个连接
 */
func GetConn(key string) (*rds, error) {
	if 	currPool, ok := rdsPoolMap[key]; ok {
		return &rds{
			connect: currPool.Get(),
		}, nil
	} else {
		return nil, errors.New("no connect")
	}
}

func GetDefaultConn() (*rds, error) {
	return GetConn(constant.RedisConfigKeyDefault)
}
/**
 * 关闭连接
 */
func (r *rds) Close() {
	_ = r.connect.Close()
}

/**
 * 获取key
 */
func buildKey(key string) string {
	return prefix + ":" + key
}

/**
 * 设置缓存
 */
func (r *rds) Set(key string, value interface{}, expire int) error {
	key = buildKey(key)
	_, err := r.connect.Do("Set", key, value, "EX", expire)
	return err
}

/**
 * key不存在时设置缓存
 */
func (r *rds) SetNX(key string, value interface{}, expire int) (bool, error) {
	key = buildKey(key)
	reply, err := r.connect.Do("Set", key, value, "EX", expire, "NX")
	//成功返回OK，失败返回nil
	return reply != nil, err
}

/**
 * 计数器加
 */
func (r *rds) Incr(key string, value interface{}, expire int) (int, error) {
	key = buildKey(key)
	reply, err := redis.Int(r.connect.Do("INCR", key))
	if err != nil {
		return 0, err
	}
	if reply == 1 {
		//初次设置的时候，设置有效期
		if _, err := r.connect.Do("EXPIRE", key, expire); err != nil {
			return 0, err
		}
	}
	return reply, nil
}

/**
 * 计数器减
 */
func (r *rds) Decr(key string) (int, error) {
	key = buildKey(key)
	reply, err := redis.Int(r.connect.Do("DECR", key))
	if err != nil {
		return 0, err
	}
	return reply, nil
}

/**
 * 设置过期时间
 */
func (r *rds) Expire(key string, expire int) error {
	key = buildKey(key)
	_, err := r.connect.Do("EXPIRE", key, expire)
	return err
}

/**
 * 获取数据
 */
func (r *rds) Get(key string) (interface{}, error) {
	key = buildKey(key)
	reply, err := r.connect.Do("Get", key)
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取整数
 */
func (r *rds) Int(key string) (int, error) {
	reply, err := redis.Int(r.Get(key))
	if err == redis.ErrNil {
		return 0, nil
	}
	return reply, err
}

/**
 * 获取整数
 */
func (r *rds) Int64(key string) (int64, error) {
	reply, err := redis.Int64(r.Get(key))
	if err == redis.ErrNil {
		return 0, nil
	}
	return reply, err
}

/**
 * 获取浮点数
 */
func (r *rds) Float64(key string) (float64, error) {
	reply, err := redis.Float64(r.Get(key))
	if err == redis.ErrNil {
		return 0, nil
	}
	return reply, err
}

/**
 * 获取字符串
 */
func (r *rds) String(key string) (string, error) {
	reply, err := redis.String(r.Get(key))
	if err == redis.ErrNil {
		return "", nil
	}
	return reply, err
}

/**
 * 获取字节切片
 */
func (r *rds) Bytes(key string) ([]byte, error) {
	reply, err := redis.Bytes(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取整数切片
 */
func (r *rds) Ints(key string) ([]int, error) {
	reply, err := redis.Ints(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取整数切片
 */
func (r *rds) Int64s(key string) ([]int64, error) {
	reply, err := redis.Int64s(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取浮点数切片
 */
func (r *rds) Float64s(key string) ([]float64, error) {
	reply, err := redis.Float64s(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取字符串切片
 */
func (r *rds) Strings(key string) ([]string, error) {
	reply, err := redis.Strings(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取字节切片的切片
 */
func (r *rds) ByteSlices(key string) ([][]byte, error) {
	reply, err := redis.ByteSlices(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取整数map
 */
func (r *rds) IntMap(key string) (map[string]int, error) {
	reply, err := redis.IntMap(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取整数map
 */
func (r *rds) Int64Map(key string) (map[string]int64, error) {
	reply, err := redis.Int64Map(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 获取字符串map
 */
func (r *rds) StringMap(key string) (map[string]string, error) {
	reply, err := redis.StringMap(r.Get(key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

/**
 * 删除缓存
 */
func (r *rds) Delete(key string) error {
	key = buildKey(key)
	_, err := r.connect.Do("DEL", key)
	return err
}

/**
 * 分布式锁
 */
func LockAndTry(lockKey string, lockFunc func() (interface{}, error), opts ...interface{}) (interface{}, error) {
	rds, _ := GetDefaultConn()
	defer rds.Close()

	wt := time.Millisecond * 100 //默认每次等待时间
	mt := 120                    //默认最大等待秒数
	for _, opt := range opts {
		switch opt.(type) {
		case int:
			mt = opt.(int) //自定义最大等待秒数
		case time.Duration:
			wt = opt.(time.Duration) //自定义每次等待时间
		}
	}

	var realWaitTime time.Duration
	for i := 1; i <= 10; i++ {
		reply, err := rds.SetNX(lockKey, 1, mt)
		if err == nil && reply {
			rst, err := lockFunc()
			_ = rds.Delete(lockKey)
			return rst, err
		}
		realWaitTime = wt * time.Duration(i)
		time.Sleep(realWaitTime)
	}

	return nil, errors.New("system exception")
}

func GetOrSet(key string, returnType string, handle func() (interface{}, error), expire int) (interface{}, error) {
	rds, err := GetDefaultConn()
	if err != nil {
		return 0, err
	}
	defer func() {
		rds.Close()
	}()
	data, err := rds.Get(key)
	if data != nil {
		var res interface{}
		switch returnType {
		case constant.TypeUint64:
			res, err = redis.Uint64(data, nil)
		default:
			res, err = redis.String(data, nil)
		}
		return res, err
	}
	data, err = handle()

	var res interface{}
	switch returnType {
	case constant.TypeUint64:
		res, err = data, err
	default:
		data2, _ := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		res, err = string(data2), nil
	}
	err = rds.Set(key, res, expire)
	return res, err
}

