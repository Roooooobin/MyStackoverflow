package answer

import (
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SelectBest(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		return
	}
	question, err := questionsdao.Find("qid = ?", answer.Qid)
	if err != nil {
		return
	}
	if question.Uid != uid {
		errMsg := errors.New("can not select a best answer if you are not the user who post the question")
		fmt.Println(errMsg)
		return
	}
	updateMap := map[string]interface{}{
		"is_best": 1,
	}
	err = answersdao.Update(updateMap, "aid = ?", aid)
	if err != nil {
		return
	}
}
