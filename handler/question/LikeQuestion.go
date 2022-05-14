package question

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionlikesdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func LikeQuestion(c *gin.Context) {

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
	qidStr := c.PostForm("qid")
	qid, err := strconv.Atoi(qidStr)
	if err != nil {
		errMsg = "Input qid error: " + err.Error()
		return
	}
	// use transaction to support concurrency and data consistency
	errTx := dao.MyDB.Transaction(func(tx *gorm.DB) error {
		questionLike := model.QuestionLike{}
		err = tx.Table(questionlikesdao.TableQuestionLikes).Where("uid = ? and qid = ?", uid, qid).
			First(&questionLike).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Already liked this question.")
		}
		questionLike = model.QuestionLike{
			Uid:  uid,
			Qid:  qid,
			Time: time.Now(),
		}
		if err = tx.Table(questionlikesdao.TableQuestionLikes).Create(questionLike).Error; err != nil {
			return err
		}
		if err = tx.Table(questionsdao.TableQuestions).Where("qid = ?", qid).
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
