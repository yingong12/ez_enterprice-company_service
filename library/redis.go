package library

import (
	"fmt"

	"github.com/go-redis/redis"
)

//自定义redis连接池
type RedisClient struct {
	*redis.Client
}
type RedisConfig struct {
	ConnectionName string //连接名称自定义
	Addr           string //地址
	Port           int    //端口
	Password       string //密码
	DB             int
	PoolSize       int //连接池大小
}

func NewRedisClient(conf *RedisConfig) (*RedisClient, error) {
	addr := fmt.Sprintf("%s:%d", conf.Addr, conf.Port)
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	if _, err := cli.Ping().Result(); err != nil {
		err = fmt.Errorf("redis client:[%s] error:%w", conf.ConnectionName, err)
		return nil, err
	}
	return &RedisClient{cli}, nil
}
