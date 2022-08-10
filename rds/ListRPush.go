package rds

func RPush(key string, value interface{}) error {

	err := redisClient.RPush(key, value).Err()
	return err
}
