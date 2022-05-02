package user

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/usersdao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	uid, ok := c.GetQuery("uid")
	if !ok {
		errMsg = "parameter uid must be passed"
	}
	user, err := usersdao.Find("uid = ?", uid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
