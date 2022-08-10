package rds

import "time"

func SetKey(key string, value interface{}, expireTime time.Duration) error {

	err := redisClient.Set(key, value, expireTime).Err()
	return err
}
