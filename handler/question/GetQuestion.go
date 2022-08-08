package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/model"
	"MyStackoverflow/rds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetQuestion(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	qid, ok := c.GetQuery("qid")
	if !ok {
		errMsg = "parameter qid must be passed"
	}
	question, err := questionsdao.Find("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	answers, err := answersdao.List("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	questionToAnswerNumMap := make(map[int]int)
	for _, answer := range answers {
		_, ok := questionToAnswerNumMap[answer.Qid]
		if !ok {
			questionToAnswerNumMap[answer.Qid] = 1
		} else {
			questionToAnswerNumMap[answer.Qid]++
		}
	}
	// attach topics
	questionToTopicsMap := make(map[int]string)
	questionTopics, err := questiontopicsdao.List("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	for _, questionTopic := range questionTopics {
		_, ok := questionToTopicsMap[questionTopic.Qid]
		topicName, _ := rds.RedisClient.Get(string(rune(questionTopic.Tid))).Result()
		if !ok {
			questionToTopicsMap[questionTopic.Qid] = topicName + ","
		} else {
			questionToTopicsMap[questionTopic.Qid] += topicName + ","
		}
	}
	numOfAnswer, ok := questionToAnswerNumMap[question.Qid]
	if !ok {
		numOfAnswer = 0
	}
	questionWithDetails := &model.QuestionWithDetails{
		Qid:         question.Qid,
		Uid:         question.Uid,
		Title:       question.Title,
		Body:        question.Body,
		Time:        question.Time,
		IsResolved:  question.IsResolved,
		Likes:       question.Likes,
		NumOfAnswer: numOfAnswer,
		Topics:      strings.TrimRight(questionToTopicsMap[question.Qid], ","),
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": questionWithDetails,
		})
	}
}
