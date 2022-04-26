package answer

import (
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RateAnswer(c *gin.Context) {

	uidStr := c.PostForm("uid")
	uid, _ := strconv.Atoi(uidStr)
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	// need to check if the question is posted by this user
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		return
	}
	question, err := questionsdao.Find("qid = ?", answer.Qid)
	if err != nil {
		return
	}
	if question.Uid != uid {
		errMsg := errors.New("can not rate the answer if you are not the user who post the question")
		fmt.Println(errMsg)
		return
	}
	ratingStr := c.PostForm("rating")
	rating, _ := strconv.Atoi(ratingStr)
	if rating < 0 || rating > 5 {
		return
	}
	updateMap := map[string]interface{}{
		"rating": rating,
	}
	err = answersdao.Update(updateMap, "aid = ?", aid)
	if err != nil {
		return
	}
}
