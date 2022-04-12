package handler

import (
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddTopic(c *gin.Context) {
	topicName := c.PostForm("topic")
	topic := model.Topic{TopicName: topicName}
	err := topicsdao.Insert(topic)
	if err != nil {
		return
	}
	t, err := topicsdao.Find("topic_name = ?", topicName)
	if err != nil {
		return
	}
	parentIdStr := c.PostForm("parent_id")
	parentId, _ := strconv.Atoi(parentIdStr)
	topicHierarchy := model.TopicHierarchy{
		Tid:       t.Tid,
		ParentTid: parentId,
	}
	err = topichierarchydao.Insert(topicHierarchy)
	if err != nil {
		return
	}
}
