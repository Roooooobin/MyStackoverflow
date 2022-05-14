package answer

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answerlikesdao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func LikeAnswer(c *gin.Context) {

	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	uidStr, ok := c.GetPostForm("uid")
	if !ok || !function.CheckNotEmpty(uidStr) {
		errMsg = "Must input uid"
		return
	}
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		errMsg = "Input uid error: " + err.Error()
		return
	}
	aidStr := c.PostForm("aid")
	aid, err := strconv.Atoi(aidStr)
	if err != nil {
		errMsg = "Input aid error: " + err.Error()
		return
	}
	// use transaction to support concurrency and data consistency
	errTx := dao.MyDB.Transaction(func(tx *gorm.DB) error {
		answerLike := model.AnswerLike{}
		err = tx.Table(answerlikesdao.TableLikes).Where("uid = ? and aid = ?", uid, aid).
			First(&answerLike).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Already liked this answer.")
		}
		answerLike = model.AnswerLike{
			Uid:  uid,
			Aid:  aid,
			Time: time.Now(),
		}
		if err = tx.Table(answerlikesdao.TableLikes).Create(answerLike).Error; err != nil {
			return err
		}
		if err = tx.Table(answersdao.TableAnswers).Where("aid = ?", aid).
			Update("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		errMsg = errTx.Error()
		return
	}
}
