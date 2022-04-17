package handler

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func AddTopic(c *gin.Context) {
	topicName := c.PostForm("topic")
	_, err := topicsdao.Find("topic_name = ?", topicName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		topic := model.Topic{TopicName: topicName}
		err = topicsdao.Insert(topic)
		if err != nil {
			return
		}
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
	cache.Init()
}
