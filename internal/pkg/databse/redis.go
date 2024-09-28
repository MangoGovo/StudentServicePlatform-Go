package database

import (
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func InitRedis() {
	addr := global.Config.GetString("redis.host")
	password := global.Config.GetString("redis.password")
	db := global.Config.GetInt("redis.DB")
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis地址
		Password: password, // 密码（如果设置了的话）
		DB:       db,       // 使用默认DB
	})
	utils.Log.Println("Redis数据库连接成功")

}
