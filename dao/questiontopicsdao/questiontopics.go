package questiontopicsdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableQuestionTopics = "QuestionTopics"

func Insert(v model.QuestionTopic) error {

	if err := dao.MyDB.Table(TableQuestionTopics).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.QuestionTopic, error) {

	questionTopic := &model.QuestionTopic{}
	res := dao.MyDB.Table(TableQuestionTopics).Where(where, values...).First(&questionTopic)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return questionTopic, nil
}

func List(where string, values ...interface{}) ([]*model.QuestionTopic, error) {

	questionTopics := make([]*model.QuestionTopic, 0)
	res := dao.MyDB.Table(TableQuestionTopics).Where(where, values...).Find(&questionTopics)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return questionTopics, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableQuestionTopics).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
