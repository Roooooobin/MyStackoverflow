package userstatsdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableUserStats = "UserStats"

func Insert(v model.UserStats) error {

	if err := dao.MyDB.Table(TableUserStats).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func InsertOrUpdate(v model.UserStats) error {

	userStats := &model.UserStats{}
	where := "uid = ?"
	uid := v.Uid
	res := dao.MyDB.Table(TableUserStats).Where(where, uid).First(&userStats)
	// new then create
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if err := dao.MyDB.Table(TableUserStats).Create(&v).Error; err != nil {
			fmt.Println("insertion failed: ", err)
			return err
		}
	} else {
		updateMap := map[string]interface{}{
			"questions_length": v.QuestionsLength + userStats.QuestionsLength,
			"questions_words":  v.QuestionsWords + userStats.QuestionsWords,
			"answers_length":   v.AnswersLength + userStats.AnswersLength,
			"answers_words":    v.AnswersWords + userStats.AnswersWords,
		}
		res = dao.MyDB.Table(TableUserStats).Where(where, uid).Updates(updateMap)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
