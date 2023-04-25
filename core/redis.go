package core

import (
	"LibraryManagement/global"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的rdb变量

func InitRedis() *redis.Client {
	return connectRedisDB(1)
}

func connectRedisDB(db int) *redis.Client {
	var confRedis = global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:               confRedis.ReturnAddr(),
		Password:           confRedis.Password,
		DB:                 db,
		PoolSize:           confRedis.PoolSize,
		IdleCheckFrequency: time.Second, // 每秒钟检查记录是否过期
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Error(fmt.Sprintf("初始化Redis数据库失败,%s", err.Error()))
		return nil
	}
	global.Log.Info(fmt.Sprintf("连接到Redis数据库成功,%s", confRedis.ReturnAddr()))
	return rdb
}
