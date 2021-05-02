package main

import (
	"fmt"
)
import "github.com/go-redis/redis"

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     RedisHost,
		Password: RedisPASS,
		DB:       0,
	})
}

func RedisSet(key string, val string) {
	redisClient.Set(key, val, 0)
}

func RedisGet(key string) string {
	return redisClient.Get(key).Val()
}

func RedisPush(key string, val ...string) {
	redisClient.RPush(key, val)
}

func RedisPop(key string) string {
	return redisClient.LPop(key).Val()
}

func RedisLen(key string) int64 {
	return redisClient.LLen(key).Val()
}

func RedisPopAll(key string) []string {
	redisLen := RedisLen(key)

	var res []string
	for i := int(redisLen); i > 0; i-- {
		temp := RedisPop(key)
		res = append(res, temp)
	}
	return res
}

func RedisTest() {
	fmt.Println("Go Redis Tutorial")

	key := "IP_LIST"
	//redisClient.RPush(key, "qqq", "www", "eee")

	lRange := redisClient.LRange(key, 0, -1)
	lLen := redisClient.LLen(key).Val()
	//pong, err := client.Ping().Result()
	fmt.Println(lRange.Val(), lLen, redisClient.LPop(key).Val())

}
