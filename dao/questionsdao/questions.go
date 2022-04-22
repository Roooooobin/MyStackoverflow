package questionsdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableQuestions = "Questions"

func Insert(v model.Question) error {

	if err := dao.MyDB.Table(TableQuestions).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.Question, error) {

	question := &model.Question{}
	res := dao.MyDB.Table(TableQuestions).Where(where, values...).First(&question)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return question, nil
}

func List(where string, values ...interface{}) ([]*model.Question, error) {

	questions := make([]*model.Question, 0)
	res := dao.MyDB.Table(TableQuestions).Where(where, values...).Find(&questions)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return questions, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableQuestions).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
