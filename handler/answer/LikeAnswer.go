package answer

import (
	"MyStackoverflow/dao/answerlikesdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func LikeAnswer(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	_, err := answerlikesdao.Find("uid = ? and aid = ?", uid, aid)
	// already added a like
	if err == nil {
		return
	}
	answerLike := model.AnswerLike{
		Uid:  uid,
		Aid:  aid,
		Time: time.Now(),
	}
	if err = answerlikesdao.Insert(answerLike); err != nil {
		// TODO: handle error
		return
	}
}
