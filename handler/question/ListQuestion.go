package question

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ListQuestion list questions with uid or qid(s)
func ListQuestion(c *gin.Context) {

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
	}
	questions := make([]*model.Question, 0)
	err := sql.Find(&questions).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": questions,
	})
}
