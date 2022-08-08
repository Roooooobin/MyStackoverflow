package rds

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

//func main() {
//
//	Init()
//	err := RedisClient.Set("1", "Backend", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//	v, err := RedisClient.Get("1").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("1", v)
//}
