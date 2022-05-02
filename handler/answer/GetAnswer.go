package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/answersdao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnswer(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	aid, ok := c.GetQuery("aid")
	if !ok {
		errMsg = "parameter aid must be passed"
	}
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": answer,
		})
	}
}
