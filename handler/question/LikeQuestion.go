package question

import (
	"MyStackoverflow/dao/questionlikesdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func LikeQuestion(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	qidStr := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidStr)
	_, err := questionlikesdao.Find("uid = ? and qid = ?", uid, qid)
	// already added a like
	if err == nil {
		return
	}
	questionLike := model.QuestionLike{
		Uid:  uid,
		Qid:  qid,
		Time: time.Now(),
	}
	if err = questionlikesdao.Insert(questionLike); err != nil {
		// TODO: handle error
		return
	}
}
