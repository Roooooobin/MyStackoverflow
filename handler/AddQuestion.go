package handler

import (
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AddQuestion(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	title := c.PostForm("title")
	body := c.PostForm("body")
	now := time.Now()
	nowFormatted := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	question := model.Question{
		Uid:   uid,
		Title: title,
		Body:  body,
		Time:  nowFormatted,
	}
	if err := questionsdao.Insert(question); err != nil {
		// TODO: handle error
		return
	}
	// get auto-generated qid
	questionJustInserted, _ := questionsdao.Find("uid = ? and time = ?", uid, nowFormatted)
	// needs to find all related topics by the hierarchy and insert into table `QuestionTopic`
	tidStr := c.PostForm("tid")
	rootTid, _ := strconv.Atoi(tidStr)
	tids := make([]int, 0)
	// TODO: will be coded into pre-computed cache later
	for rootTid != 0 {
		tids = append(tids, rootTid)
		topic, _ := topichierarchydao.Find("tid = ?", rootTid)
		if topic != nil {
			rootTid = topic.ParentTid
		}
	}
	for _, tid := range tids {
		//fmt.Println(tid)
		questionTopic := model.QuestionTopic{
			Qid: questionJustInserted.Qid,
			Tid: tid,
		}
		err := questiontopicsdao.Insert(questionTopic)
		if err != nil {
			return
		}
	}
}
