package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answerlikesdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func LikeAnswer(c *gin.Context) {

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
	aidStr := c.PostForm("aid")
	aid, err := strconv.Atoi(aidStr)
	if err != nil {
		errMsg = "Input aid error: " + err.Error()
		return
	}
	_, err = answerlikesdao.Find("uid = ? and aid = ?", uid, aid)
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
		errMsg = err.Error()
		return
	}
}
