package user

import (
	"MyStackoverflow/dao/usersdao"
	"github.com/gin-gonic/gin"
	"strings"
)

func EditUser(c *gin.Context) {

	uid, ok := c.GetPostForm("uid")
	if !ok {
		// TODO: error handling
		return
	}
	updateMap := make(map[string]interface{})
	checkNotEmpty := func(s string) bool {
		return strings.Trim(s, " ") != ""
	}
	username := c.PostForm("username")
	if checkNotEmpty(username) {
		updateMap["username"] = username
	}
	email := c.PostForm("email")
	if checkNotEmpty(email) {
		updateMap["email"] = email
	}
	password := c.PostForm("password")
	if checkNotEmpty(password) {
		updateMap["password"] = password
	}
	city := c.PostForm("city")
	if checkNotEmpty(city) {
		updateMap["city"] = city
	}
	state := c.PostForm("state")
	if checkNotEmpty(state) {
		updateMap["state"] = state
	}
	country := c.PostForm("country")
	if checkNotEmpty(country) {
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
