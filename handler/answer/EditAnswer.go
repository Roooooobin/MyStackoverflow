package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
	"strconv"
)

func EditAnswer(c *gin.Context) {

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
	// need to check if the question is posted by this user
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if answer.Uid != uid {
		errMsg = "You can not edit answer that is posted by others!"
		return
	}
	body := c.PostForm("body")
	updateMap := map[string]interface{}{
		"body": body,
	}
	err = answersdao.Update(updateMap, "aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
}
