package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/questionsdao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuestion(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	qid, ok := c.GetQuery("qid")
	if !ok {
		errMsg = "parameter qid must be passed"
	}
	question, err := questionsdao.Find("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": question,
		})
	}
}
