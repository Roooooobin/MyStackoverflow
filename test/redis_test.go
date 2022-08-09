package main

import (
	"MyStackoverflow/rds"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	rds.Init()
	rds.RedisClient.RPush("1000", 1)
	rds.RedisClient.RPush("1000", 2)
	rds.RedisClient.LPush("1000", 3)
	vals, _ := rds.RedisClient.LRange("1000", 0, -1).Result()
	fmt.Println(vals)
}
