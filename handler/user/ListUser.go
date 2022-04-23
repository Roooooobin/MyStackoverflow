package user

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(c *gin.Context) {

	sql := dao.MyDB.Table(usersdao.TableUsers)
	uid, ok := c.GetQuery("uid")
	if ok {
		sql.Where("uid = ?", uid)
	}
	users := make([]*model.User, 0)
	err := sql.Find(&users).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
