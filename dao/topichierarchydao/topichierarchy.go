package topichierarchydao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var tableTopicHierarchy = "TopicHierarchy"

func Insert(v model.Topic) error {

	if err := dao.MyDB.Table(tableTopicHierarchy).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.TopicHierarchy, error) {

	t := &model.TopicHierarchy{}
	res := dao.MyDB.Table(tableTopicHierarchy).Where(where, values...).First(&t)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return t, nil
}

func List(where string, values ...interface{}) ([]*model.TopicHierarchy, error) {

	topics := make([]*model.TopicHierarchy, 0)
	res := dao.MyDB.Table(tableTopicHierarchy).Where(where, values...).Find(&topics)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return topics, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(tableTopicHierarchy).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
