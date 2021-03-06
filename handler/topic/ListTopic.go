package topic

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTopic(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	topics := make([]*model.Topic, 0)
	err := dao.MyDB.Table(topicsdao.TableTopics).Order("tid asc").Find(&topics).Error
	if err != nil {
		errMsg = err.Error()
		return
	}
	topicsReturn := make([]map[string]interface{}, 0)
	for _, topic := range topics {
		topicsReturn = append(topicsReturn, map[string]interface{}{
			"value": topic.Tid,
			"label": topic.TopicName,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": topicsReturn,
	})
}
