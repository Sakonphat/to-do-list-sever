package config

import (
	"fmt"
	"sever/utils"
)

type RedisConfig struct {
	Host     string
	Port     string
}

func GetRedisConfig() *RedisConfig {

	redisConfig := RedisConfig{
		Host: utils.Env("REDIS_HOST"),
		Port: utils.Env("REDIS_PORT"),
	}
	
	return &redisConfig
}

func GetRedisDns(redisConfig *RedisConfig) string {
	return fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)
}
