package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
	"strconv"
)

func MarkResolved(c *gin.Context) {

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
		errMsg = "Can not select a best answer if you are not the user who post the question!"
		return
	}
	isResolved := 1
	// already is best, cancel it
	if question.IsResolved == 1 {
		isResolved = 0
	}
	updateMap := map[string]interface{}{
		"is_resolved": isResolved,
	}
	err = questionsdao.Update(updateMap, "qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
}
