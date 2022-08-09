package answertopicsdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableAnswerTopics = "AnswerTopics"

func Insert(v model.AnswerTopic) error {

	if err := dao.MyDB.Table(TableAnswerTopics).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.AnswerTopic, error) {

	answerTopic := &model.AnswerTopic{}
	res := dao.MyDB.Table(TableAnswerTopics).Where(where, values...).First(&answerTopic)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return answerTopic, nil
}

func List(where string, values ...interface{}) ([]*model.AnswerTopic, error) {

	answerTopics := make([]*model.AnswerTopic, 0)
	res := dao.MyDB.Table(TableAnswerTopics).Where(where, values...).Find(&answerTopics)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return answerTopics, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableAnswerTopics).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
