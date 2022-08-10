package rds

func GetListValues(key string) ([]string, error) {

	values, err := redisClient.LRange(key, 0, -1).Result()
	return values, err
}
