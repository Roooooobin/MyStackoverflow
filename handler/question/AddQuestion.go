package question

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func AddQuestion(c *gin.Context) {

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
	title := c.PostForm("title")
	body := c.PostForm("body")
	now := time.Now()
	nowFormatted := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	question := &model.Question{
		Qid:   0,
		Uid:   uid,
		Title: title,
		Body:  body,
		Time:  nowFormatted,
	}
	errTx := dao.MyDB.Transaction(func(tx *gorm.DB) error {

		if err = tx.Table(questionsdao.TableQuestions).Create(question).Error; err != nil {
			// return any error will rollback
			return err
		}
		qid := question.Qid
		// get auto-generated qid
		//questionJustInserted, _ := questionsdao.Find("uid = ? and time = ?", uid, nowFormatted)
		// needs to find all related topics by the hierarchy and insert into table `QuestionTopic`
		tidStr := c.PostForm("tid")
		rootTid, errR := strconv.Atoi(tidStr)
		if errR != nil {
			return errR
		}
		tids := cache.ParentTopics[rootTid]
		for _, tid := range tids {
			//fmt.Println(tid)
			questionTopic := &model.QuestionTopic{
				Qid: qid,
				Tid: tid,
			}
			err = tx.Table(questiontopicsdao.TableQuestionTopics).Create(&questionTopic).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if errTx != nil {
		errMsg = errTx.Error()
		return
	}
}
