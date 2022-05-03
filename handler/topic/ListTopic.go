package topic

import (
	"MyStackoverflow/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTopic(c *gin.Context) {

	topicName2SubTopicNames := make(map[string][]string)
	topic2SubTopics := cache.Topic2SubTopics
	topicID2Name := cache.TopicID2Name
	for ptid, tids := range topic2SubTopics {
		names := make([]string, 0)
		for _, tid := range tids {
			names = append(names, topicID2Name[tid])
		}
		topicName2SubTopicNames[topicID2Name[ptid]] = names
	}
	c.JSON(http.StatusOK, gin.H{
		"data": topicName2SubTopicNames,
	})
}
