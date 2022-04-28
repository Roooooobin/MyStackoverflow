package user

import (
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/function"
	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {

	uid, ok := c.GetPostForm("uid")
	if !ok {
		// TODO: error handling
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
	err := usersdao.Update(updateMap, "uid = ?", uid)
	if err != nil {
		return
	}
}
