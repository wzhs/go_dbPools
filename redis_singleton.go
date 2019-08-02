package go_dbPools

import (
	"sync"
	"go_dbPools/redis"
)

var once sync.Once
var redisman *RedisMan

type RedisMan struct {
	Pool *redis.ConnPool
}

func GetRedisInstance() *RedisMan {
	once.Do(func() {
		redisman = NewInstance()
	})
	return redisman
}

/*
host, password string, database, maxOpenConns, maxIdleConns int
*/
var _host string
var _pass string
var _db int
var _maxOpens int
var _maxIdels int

func SetRedisParas(host, password string, database, maxOpenConns, maxIdleConns int) {
	_host = host
	_pass = password
	_db = database
	_maxOpens = maxOpenConns
	_maxIdels = maxIdleConns

}

func NewInstance() *RedisMan {
	red := &RedisMan{
		Pool: redis.InitRedisPool(_host, _pass, _db, _maxOpens, _maxIdels),
	}
	//默认连接，没有设置参数的时候
	if len(_host) == 0 {
		red.Pool = redis.InitRedisPool("222.186.136.51:7910", "shiyuxing521", 0, -1, -1)
	}
	return red
}
