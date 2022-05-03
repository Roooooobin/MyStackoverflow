package cache

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
)

var TopicID2Name map[int]string

func GetTopicNameByID() map[int]string {

	topicID2Name := make(map[int]string)
	sql := dao.MyDB.Table(topicsdao.TableTopics)
	allTopics := make([]*model.Topic, 0)
	sql.Find(&allTopics)
	for _, topic := range allTopics {
		topicID2Name[topic.Tid] = topic.TopicName
	}
	return topicID2Name
}
