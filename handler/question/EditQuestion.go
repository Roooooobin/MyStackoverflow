package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
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
