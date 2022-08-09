package util

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", viper.GetString("redis.address"))
		}}
}

func GetRedisConn() redis.Conn {
	conn := pool.Get()
	return conn
}
