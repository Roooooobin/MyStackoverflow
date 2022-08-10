package rds

func GetValue(key string) (string, error) {

	res, err := redisClient.Get(key).Result()
	return res, err
}
