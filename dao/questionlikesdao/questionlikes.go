package questionlikesdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableQuestionLikes = "QuestionLikes"

func Insert(v model.QuestionLike) error {

	if err := dao.MyDB.Table(TableQuestionLikes).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.QuestionLike, error) {

	questionLike := &model.QuestionLike{}
	res := dao.MyDB.Table(TableQuestionLikes).Where(where, values...).First(&questionLike)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return questionLike, nil
}

func List(where string, values ...interface{}) ([]*model.QuestionLike, error) {

	questionLikes := make([]*model.QuestionLike, 0)
	res := dao.MyDB.Table(TableQuestionLikes).Where(where, values...).Find(&questionLikes)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return questionLikes, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableQuestionLikes).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
