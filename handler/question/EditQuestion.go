package question

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func EditQuestion(c *gin.Context) {

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
	qidStr, ok := c.GetPostForm("qid")
	if !ok || !function.CheckNotEmpty(qidStr) {
		errMsg = "Must input qid"
		return
	}
	qid, err := strconv.Atoi(qidStr)
	if err != nil {
		errMsg = "Input qid error: " + err.Error()
		return
	}
	question, err := questionsdao.Find("qid = ?", qid)
	if err != nil {
		errMsg = err.Error()
		return
	}
	if question.Uid != uid {
		errMsg = "You can not edit question that is posted by others!"
		return
	}
	updateMap := make(map[string]interface{})
	title, ok := c.GetPostForm("title")
	if ok {
		updateMap["title"] = title
	}
	body, ok := c.GetPostForm("body")
	if ok {
		updateMap["body"] = body
	}
	_, ok = c.GetPostForm("isResolved")
	if ok {
		updateMap["is_resolved"] = 1
	}
	// modify tid
	tidStr := c.PostForm("tid")
	// if topics of the question need to be modified
	if function.CheckNotEmpty(tidStr) {
		errTx := dao.MyDB.Transaction(func(tx *gorm.DB) error {
			// delete first and then add new topics
			err := tx.Table(questiontopicsdao.TableQuestionTopics).Where("qid = ?", qid).
				Delete(&model.QuestionTopic{}).Error
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
				questionTopic := &model.QuestionTopic{
					Qid: qid,
					Tid: tid,
				}
				err = tx.Table(questiontopicsdao.TableQuestionTopics).Create(&questionTopic).Error
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
		if err = questionsdao.Update(updateMap, "qid = ?", qid); err != nil {
			errMsg = err.Error()
			return
		}
	}
}
