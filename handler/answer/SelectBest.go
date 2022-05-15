package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SelectBest(c *gin.Context) {

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
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	question, err := questionsdao.Find("qid = ?", answer.Qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if question.Uid != uid {
		errMsg = "Can not select a best answer if you are not the user who post the question!"
		return
	}
	is_best := 1
	// already is best, cancel it
	if answer.IsBest == 1 {
		is_best = 0
	}
	updateMap := map[string]interface{}{
		"is_best": is_best,
	}
	err = answersdao.Update(updateMap, "aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
}
