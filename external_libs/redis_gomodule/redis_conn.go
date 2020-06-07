package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	MaxActive      int = 400
	MaxIdle            = 20
	MaxIdleTimeout     = 240
	Timeout            = 30
)

//redis 连接需要的配置信息
type ConfigST struct {
	User string `json:"User" yaml:"User"`
	Pw   string `json:"Pw" yaml:"Pw"`
	Host string `json:"Host" yaml:"Host"`
	Db   int    `json:"Db" yaml:"Db"`
}

var redisClient *redis.Pool

func InitRedis(cfg ConfigST) error {
	redisClient = &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: MaxIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cfg.Host,
				redis.DialPassword(cfg.Pw),
				redis.DialDatabase(cfg.Db),
				redis.DialConnectTimeout(Timeout*time.Second),
				redis.DialReadTimeout(Timeout*time.Second),
				redis.DialWriteTimeout(Timeout*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	// 从池里获取连接
	rc := redisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	// 错误判断
	return rc.Err()
}

func WriteMapDataWithExpire(data map[string]interface{}, expire int64) (err error) {
	rc := redisClient.Get()
	defer rc.Close()
	for k, v := range data {
		if _, err = rc.Do("Set", k, v); err != nil {
			return err
		}
		if _, err = rc.Do("EXPIRE", k, expire); err != nil {
			return err
		}
	}
	return nil
}

func GetKeyString(key string) (string, error) {
	rc := redisClient.Get()
	defer rc.Close()
	val, err := redis.String(rc.Do("Get", key))
	return val, err
}
