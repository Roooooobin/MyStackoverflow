package rds

func Init() {

	if RedisClient == nil {
		RedisClient = NewRedisClient()
	}
	GetTopicNameByID()
	GetParentTopics()
}
