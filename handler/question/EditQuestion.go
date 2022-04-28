package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/questionsdao"
	"github.com/gin-gonic/gin"
	"strconv"
)

func EditQuestion(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	qidStr := c.PostForm("qid")
	qid, _ := strconv.Atoi(qidStr)
	question, err := questionsdao.Find("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if question.Uid != uid {
		errMsg = "You can not edit question that is posted by others!"
		return
	}
	updateMap := make(map[string]interface{})
	title, ok := c.GetPostForm("title")
	if ok {
		updateMap["title"] = title
	}
	body, ok := c.GetPostForm("body")
	if ok {
		updateMap["body"] = body
	}
	_, ok = c.GetPostForm("isResolved")
	if ok {
		updateMap["is_resolved"] = 1
	}
	if err = questionsdao.Update(updateMap, "qid = ?", qid); err != nil {
		errMsg = err.Error()
		return
	}
}
