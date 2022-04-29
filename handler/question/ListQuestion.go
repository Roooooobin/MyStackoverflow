package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ListQuestion list questions by uid or qid(s)
func ListQuestion(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	sql := dao.MyDB.Table(questionsdao.TableQuestions)
	uid, ok := c.GetQuery("uid")
	if ok {
		sql.Where("uid = ?", uid)
	}
	qid, ok := c.GetQuery("qid")
	if ok {
		qidList := strings.Split(qid, ",")
		sql.Where("qid in (?)", qidList)
	}
	sortByLikes, ok := c.GetQuery("sortByLikes")
	if ok {
		sql.Order("likes " + sortByLikes)
	}
	sortByTime, ok := c.GetQuery("sortByTime")
	if ok {
		sql.Order("time " + sortByTime)
	} else {
		// list in reverse chronological order by default
		sql.Order("time desc")
	}
	questions := make([]*model.Question, 0)
	err := sql.Find(&questions).Error
	if err != nil {
		errMsg = err.Error()
		return
	}
	// filter questions only within the topic
	tid, ok := c.GetQuery("tid")
	if ok {
		questionTopics, err := questiontopicsdao.List("tid = ?", tid)
		if err != nil {
			errMsg = err.Error()
			return
		}
		qidSet := make(map[int]struct{})
		for _, qt := range questionTopics {
			qidSet[qt.Qid] = struct{}{}
		}
		tmp := make([]*model.Question, 0)
		for _, question := range questions {
			if _, ok := qidSet[question.Qid]; ok {
				tmp = append(tmp, question)
			}
		}
		questions = tmp
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": questions,
		})
	}
}
