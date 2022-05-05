package user

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	uid, ok := c.GetPostForm("uid")
	if !ok || !function.CheckNotEmpty(uid) {
		errMsg = "Must input uid"
		return
	}
	updateMap := make(map[string]interface{})
	username := c.PostForm("username")
	if function.CheckNotEmpty(username) {
		updateMap["username"] = username
	}
	email := c.PostForm("email")
	if function.CheckNotEmpty(email) {
		updateMap["email"] = email
	}
	password := c.PostForm("password")
	if function.CheckNotEmpty(password) {
		updateMap["password"] = password
	}
	city := c.PostForm("city")
	if function.CheckNotEmpty(city) {
		updateMap["city"] = city
	}
	state := c.PostForm("state")
	if function.CheckNotEmpty(state) {
		updateMap["state"] = state
	}
	country := c.PostForm("country")
	if function.CheckNotEmpty(country) {
		updateMap["country"] = country
	}
	profile, ok := c.GetPostForm("profile")
	if ok {
		updateMap["profile"] = profile
	}
	if len(updateMap) > 0 {
		err := usersdao.Update(updateMap, "uid = ?", uid)
		if err != nil {
			errMsg = err.Error()
			return
		}
	}
}
