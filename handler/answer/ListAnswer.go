package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/model"
	"MyStackoverflow/rds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ListAnswer list answer with aid(s) or qid(s) or uid
func ListAnswer(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	sql := dao.MyDB.Table(answersdao.TableAnswers)
	uid, ok := c.GetQuery("uid")
	if ok {
		sql.Where("uid = ?", uid)
	}
	qid, ok := c.GetQuery("qid")
	if ok {
		qidList := strings.Split(qid, ",")
		sql.Where("qid in (?)", qidList)
	}
	aid, ok := c.GetQuery("aid")
	if ok {
		aidList := strings.Split(aid, ",")
		sql.Where("aid in (?)", aidList)
	}
	// sort by (time / likes)
	sortByLikes, ok := c.GetQuery("sortByLikes")
	if ok {
		sql.Order("likes " + sortByLikes)
	}
	sortByTime, ok := c.GetQuery("sortByTime")
	if ok {
		sql.Order("time " + sortByTime)
	} else {
		sql.Order("time desc")
	}
	answers := make([]*model.Answer, 0)
	err := sql.Find(&answers).Error
	if err != nil {
		errMsg = err.Error()
		return
	}
	aids := make([]int, 0)
	for _, answer := range answers {
		aids = append(aids, answer.Aid)
	}
	// attach topics
	answerToTopicsMap := make(map[int]string)
	answerTopics, err := answertopicsdao.List("aid in (?)", aids)
	if err != nil {
		errMsg = err.Error()
		return
	}
	for _, answerTopic := range answerTopics {
		_, ok := answerToTopicsMap[answerTopic.Aid]
		topicName, _ := rds.RedisClient.Get(string(rune(answerTopic.Tid))).Result()
		if !ok {
			answerToTopicsMap[answerTopic.Aid] = topicName + ","
		} else {
			answerToTopicsMap[answerTopic.Aid] += topicName + ","
		}
	}
	answerWithDetails := make([]*model.AnswerWithDetails, 0)
	for _, answer := range answers {
		answerWithDetails = append(answerWithDetails, &model.AnswerWithDetails{
			Aid:    answer.Aid,
			Qid:    answer.Qid,
			Uid:    answer.Uid,
			Body:   answer.Body,
			Time:   answer.Time,
			IsBest: answer.IsBest,
			Likes:  answer.Likes,
			Rating: answer.Rating,
			Topics: strings.TrimRight(answerToTopicsMap[answer.Aid], ","),
		})
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": answerWithDetails,
		})
	}
}
