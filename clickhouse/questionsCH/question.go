package questionsCH

import (
	"MyStackoverflow/clickhouse"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableQuestions = "Questions"

func Transform(v *model.Question) model.QuestionCH {

	return model.QuestionCH{
		Qid:        int64(v.Qid),
		Uid:        int64(v.Uid),
		Title:      v.Title,
		Body:       v.Body,
		Time:       v.Time,
		IsResolved: int8(v.IsResolved),
		Likes:      int32(v.Likes),
	}
}

func Insert(v model.QuestionCH) error {

	if err := clickhouse.ClickHouseDB.Table(TableQuestions).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.QuestionCH, error) {

	question := &model.QuestionCH{}
	res := clickhouse.ClickHouseDB.Table(TableQuestions).Where(where, values...).First(&question)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return question, nil
}

func List(where string, values ...interface{}) ([]*model.QuestionCH, error) {

	questions := make([]*model.QuestionCH, 0)
	res := clickhouse.ClickHouseDB.Table(TableQuestions).Where(where, values...).Find(&questions)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return questions, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := clickhouse.ClickHouseDB.Table(TableQuestions).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
