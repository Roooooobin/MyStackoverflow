package rds

func DeleteKey(key string) error {

	err := redisClient.Del(key).Err()
	return err
}
