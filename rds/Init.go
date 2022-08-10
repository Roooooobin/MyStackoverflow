package rds

import "github.com/go-redis/redis"

var redisClient *redis.Client

func Init() {

	if redisClient == nil {
		redisClient = newRedisClient()
	}
	GetTopicNameByID()
	GetParentTopics()
}

func newRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}
