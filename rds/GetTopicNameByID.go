package rds

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
	"strconv"
)

func GetTopicNameByID() {

	sql := dao.MyDB.Table(topicsdao.TableTopics)
	allTopics := make([]*model.Topic, 0)
	sql.Find(&allTopics)
	for _, topic := range allTopics {
		RedisClient.Set(strconv.Itoa(topic.Tid), topic.TopicName, 0)
	}
}
