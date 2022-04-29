package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/model"
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
	}
	answers := make([]*model.Answer, 0)
	err := sql.Find(&answers).Error
	if err != nil {
		errMsg = err.Error()
		return
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": answers,
		})
	}
}
