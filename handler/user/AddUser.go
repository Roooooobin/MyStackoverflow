package user

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	city := c.PostForm("city")
	state := c.PostForm("state")
	country := c.PostForm("country")
	profile := c.PostForm("profile")
	user := model.User{
		Username: username,
		Email:    email,
		Password: password,
		City:     city,
		State:    state,
		Country:  country,
		Profile:  profile,
	}
	err := usersdao.Insert(user)
	if err != nil {
		return
	}
}
