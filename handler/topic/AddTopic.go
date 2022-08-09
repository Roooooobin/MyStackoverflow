package topic

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
	"MyStackoverflow/rds"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func AddTopic(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	topicName := c.PostForm("topic")
	_, err := topicsdao.Find("topic_name = ?", topicName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		topic := model.Topic{TopicName: topicName}
		err = topicsdao.Insert(topic)
		if err != nil {
			errMsg = err.Error()
			return
		}
	}
	t, err := topicsdao.Find("topic_name = ?", topicName)
	if err != nil {
		errMsg = err.Error()
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
		errMsg = err.Error()
		return
	}
	rds.Init()
}
