package answersdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableAnswers = "Answers"

func Insert(v model.Answer) error {

	if err := dao.MyDB.Table(TableAnswers).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.Answer, error) {

	answer := &model.Answer{}
	res := dao.MyDB.Table(TableAnswers).Where(where, values...).First(&answer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return answer, nil
}

func List(where string, values ...interface{}) ([]*model.Answer, error) {

	answers := make([]*model.Answer, 0)
	res := dao.MyDB.Table(TableAnswers).Where(where, values...).Find(&answers)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return answers, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableAnswers).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
