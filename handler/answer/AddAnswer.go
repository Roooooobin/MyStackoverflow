package answer

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func AddAnswer(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	qidStr := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidStr)
	body := c.PostForm("body")
	now := time.Now()
	nowFormatted := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	tidStr := c.PostForm("tid")
	rootTid, _ := strconv.Atoi(tidStr)
	// needs to find all related topics by the hierarchy and insert into table `AnswerTopic`
	tids := cache.ParentTopics[rootTid]
	answer := &model.Answer{
		Uid:  uid,
		Qid:  qid,
		Body: body,
		Time: nowFormatted,
	}
	if err := dao.MyDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(answersdao.TableAnswers).Create(answer).Error; err != nil {
			return err
		}
		// get auto-generated qid
		aid := answer.Aid
		for _, tid := range tids {
			answerTopic := &model.AnswerTopic{
				Aid: aid,
				Tid: tid,
			}
			err := tx.Table(answertopicsdao.TableAnswerTopics).Create(answerTopic).Error
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		// TODO: handle error
		return
	}
}
