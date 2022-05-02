package user

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ListUser(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	sql := dao.MyDB.Table(usersdao.TableUsers)
	uidStr, ok := c.GetQuery("uid")
	if ok {
		uids := strings.Split(uidStr, ",")
		sql.Where("uid in (?)", uids)
	}
	users := make([]*model.User, 0)
	err := sql.Find(&users).Error
	if err != nil {
		errMsg = err.Error()
		return
	}
	// should not(no need either) return the password
	for _, user := range users {
		user.Password = ""
	}
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	}

}
