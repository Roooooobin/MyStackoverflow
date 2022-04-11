package answerlikesdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var tableLikes = "AnswerLikes"

func Insert(v model.AnswerLike) error {

	if err := dao.MyDB.Table(tableLikes).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.AnswerLike, error) {

	answerLike := &model.AnswerLike{}
	res := dao.MyDB.Table(tableLikes).Where(where, values...).First(&answerLike)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return answerLike, nil
}

func List(where string, values ...interface{}) ([]*model.AnswerLike, error) {

	answerLikes := make([]*model.AnswerLike, 0)
	res := dao.MyDB.Table(tableLikes).Where(where, values...).Find(&answerLikes)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return answerLikes, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(tableLikes).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
