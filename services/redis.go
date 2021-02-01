package services

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"sever/config"
	"time"
)

var client *redis.Client

func SetupRedis() {

	dns := config.GetRedisDns(config.GetRedisConfig())
	fmt.Printf("REDIS PORT : %s \n", dns)

	client = redis.NewClient(&redis.Options{
		Addr: dns,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func StoreRedis(uuid string, tokenDetails *TokenDetails ) error {
	accessTime := time.Unix(tokenDetails.AtExpires, 0)
	now := time.Now()

	err := client.Set(tokenDetails.AccessUuid, uuid, accessTime.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}
