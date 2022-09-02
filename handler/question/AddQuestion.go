package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"MyStackoverflow/mq"
	"MyStackoverflow/rds"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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
		// needs to find all related topics by the hierarchy and insert into table `QuestionTopic`
		tidStr := c.PostForm("tid")
		// if the question is not assigned to any topic, return
		if !function.CheckNotEmpty(tidStr) {
			return nil
		}
		rootTid, errR := strconv.Atoi(tidStr)
		if errR != nil {
			return errR
		}
		key := rds.FormParentsKey(rootTid)
		tids, _ := rds.GetListValues(key)
		for _, ttid := range tids {
			tid, _ := strconv.Atoi(ttid)
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
	} else {
		// add the question to es
		//err = es.AddQuestion(question)
		//if err != nil {
		//	errMsg = err.Error()
		//	return
		//}
		// add the question to clickhouse
		//err = questionsCH.Insert(questionsCH.Transform(question))
		//if err != nil {
		//	errMsg = err.Error()
		//	return
		//}
		// add message to mq
		msg := &primitive.Message{
			Topic: "questions",
			Body:  []byte(strconv.Itoa(question.Uid) + ":" + question.Body),
		}
		// 发送信息
		err = mq.P.SendOneWay(context.Background(), msg)
		if err != nil {
			fmt.Printf("send message error:%s\n", err)
		} else {
			fmt.Printf("send message success\n")
		}
	}
}
