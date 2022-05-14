package user

import (
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeUser(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(http.StatusOK, gin.H{
				"data": -1,
			})
		}
	}()
	username := c.PostForm("username")
	if !function.CheckNotEmpty(username) {
		errMsg = "must input username"
		return
	}
	password := c.PostForm("password")
	if !function.CheckNotEmpty(password) {
		errMsg = "must input password"
		return
	}
	user, err := usersdao.Find("username = ?", username)
	if err != nil || password != user.Password {
		errMsg = "authorization not passed"
		return
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": user.Uid,
		})
	}
}
