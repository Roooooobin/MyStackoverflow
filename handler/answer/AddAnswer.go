package answer

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
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
	answer := model.Answer{
		Uid:  uid,
		Qid:  qid,
		Body: body,
		Time: nowFormatted,
	}
	if err := answersdao.Insert(answer); err != nil {
		// TODO: handle error
		return
	}
	// get auto-generated qid
	answerJustInserted, _ := answersdao.Find("uid = ? and qid = ? and time = ?", uid, qid, nowFormatted)
	// needs to find all related topics by the hierarchy and insert into table `QuestionTopic`
	tidStr := c.PostForm("tid")
	rootTid, _ := strconv.Atoi(tidStr)
	tids := cache.ParentTopics[rootTid]
	for _, tid := range tids {
		//fmt.Println(tid)
		answerTopic := model.AnswerTopic{
			Aid: answerJustInserted.Aid,
			Tid: tid,
		}
		err := answertopicsdao.Insert(answerTopic)
		if err != nil {
			return
		}
	}
}
