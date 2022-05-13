package answer

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func EditAnswer(c *gin.Context) {

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
	// need to check if the question is posted by this user
	answer, err := answersdao.Find("aid = ?", aid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if answer.Uid != uid {
		errMsg = "You can not edit answer that is posted by others!"
		return
	}
	body := c.PostForm("body")
	updateMap := map[string]interface{}{
		"body": body,
	}
	// modify tid
	tidStr := c.PostForm("tid")
	// if topics of the question need to be modified
	if function.CheckNotEmpty(tidStr) {
		errTx := dao.MyDB.Transaction(func(tx *gorm.DB) error {
			// delete first and then add new topics
			err := tx.Table(answertopicsdao.TableAnswerTopics).Where("aid = ?", aid).
				Delete(&model.AnswerTopic{}).Error
			if err != nil {
				return err
			}
			// needs to find all related topics by the hierarchy and insert into table `QuestionTopic`
			rootTid, errR := strconv.Atoi(tidStr)
			if errR != nil {
				return errR
			}
			tids := cache.ParentTopics[rootTid]
			for _, tid := range tids {
				answerTopic := &model.AnswerTopic{
					Aid: aid,
					Tid: tid,
				}
				err = tx.Table(answertopicsdao.TableAnswerTopics).Create(&answerTopic).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
		if errTx != nil {
			errMsg = errTx.Error()
			return
		}
	}
	if len(updateMap) > 0 {
		err = answersdao.Update(updateMap, "aid = ?", aid)
		if err != nil {
			errMsg = err.Error()
			return
		}
	}
}
