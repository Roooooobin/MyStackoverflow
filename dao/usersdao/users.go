package usersdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var TableUsers = "Users"

func Insert(v model.User) error {

	// with password md5 encrypted to be safer(deprecated, md5 by frontend)
	//hash := md5.Sum([]byte(v.Password))
	//v.Password = hex.EncodeToString(hash[:])
	if err := dao.MyDB.Table(TableUsers).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.User, error) {

	user := &model.User{}
	res := dao.MyDB.Table(TableUsers).Where(where, values...).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return user, nil
}

func List(where string, values ...interface{}) ([]*model.User, error) {

	users := make([]*model.User, 0)
	res := dao.MyDB.Table(TableUsers).Where(where, values...).Find(&users)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return users, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(TableUsers).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
