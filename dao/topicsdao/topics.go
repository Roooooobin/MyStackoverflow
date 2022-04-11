package topicsdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var tableTopics = "Topics"

func Insert(v model.Topic) error {

	if err := dao.MyDB.Table(tableTopics).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.Topic, error) {

	topic := &model.Topic{}
	res := dao.MyDB.Table(tableTopics).Where(where, values...).First(&topic)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return topic, nil
}

func List(where string, values ...interface{}) ([]*model.Topic, error) {

	topics := make([]*model.Topic, 0)
	res := dao.MyDB.Table(tableTopics).Where(where, values...).Find(&topics)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return topics, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(tableTopics).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
