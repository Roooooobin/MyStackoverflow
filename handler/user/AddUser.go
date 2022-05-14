package user

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
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
	err := dao.MyDB.Table(usersdao.TableUsers).Create(&user).Error
	if err != nil {
		errMsg = err.Error()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": user.Uid,
		})
	}
}
