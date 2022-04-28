package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answersdao"
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
	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
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
