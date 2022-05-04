package answer

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func AddAnswer(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	uidStr, ok := c.GetPostForm("uid")
	if !ok || !function.CheckNotEmpty(uidStr) {
		errMsg = "Must input uid"
		return
	}
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		errMsg = "Input uid error: " + err.Error()
		return
	}
	qidStr := c.PostForm("qid")
	qid, err := strconv.Atoi(qidStr)
	if err != nil {
		errMsg = "input qid error: " + err.Error()
		return
	}
	body := c.PostForm("body")
	now := time.Now()
	nowFormatted := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	answer := &model.Answer{
		Uid:  uid,
		Qid:  qid,
		Body: body,
		Time: nowFormatted,
	}
	if err = dao.MyDB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(answersdao.TableAnswers).Create(answer).Error; err != nil {
			return err
		}
		// get auto-generated qid
		aid := answer.Aid
		tidStr := c.PostForm("tid")
		// if the answer is not assigned to any topic, return
		if !function.CheckNotEmpty(tidStr) {
			return nil
		}
		rootTid, err := strconv.Atoi(tidStr)
		if err != nil {
			return err
		}
		// needs to find all related topics by the hierarchy and insert into table `AnswerTopic`
		tids := cache.ParentTopics[rootTid]
		for _, tid := range tids {
			answerTopic := &model.AnswerTopic{
				Aid: aid,
				Tid: tid,
			}
			err = tx.Table(answertopicsdao.TableAnswerTopics).Create(answerTopic).Error
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		errMsg = "transaction: failed" + err.Error()
		return
	}
}
