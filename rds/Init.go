package rds

func Init() {

	RedisClient = NewRedisClient()
	GetTopicNameByID()
}
