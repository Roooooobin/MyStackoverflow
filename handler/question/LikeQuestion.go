package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/questionlikesdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func LikeQuestion(c *gin.Context) {

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
		errMsg = "Input qid error: " + err.Error()
		return
	}
	_, err = questionlikesdao.Find("uid = ? and qid = ?", uid, qid)
	// already added a like
	if err == nil {
		errMsg = "Already liked this question."
		return
	}
	questionLike := model.QuestionLike{
		Uid:  uid,
		Qid:  qid,
		Time: time.Now(),
	}
	if err = questionlikesdao.Insert(questionLike); err != nil {
		errMsg = err.Error()
		return
	}
}
