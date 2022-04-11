package usersdao

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var tableUsers = "Users"

func Insert(v model.User) error {

	// with password md5 encrypted to be safer
	hash := md5.Sum([]byte(v.Password))
	v.Password = hex.EncodeToString(hash[:])
	if err := dao.MyDB.Table(tableUsers).Create(&v).Error; err != nil {
		fmt.Println("insertion failed: ", err)
		return err
	}
	return nil
}

func Find(where string, values ...interface{}) (*model.User, error) {

	user := &model.User{}
	res := dao.MyDB.Table(tableUsers).Where(where, values...).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return user, nil
}

func List(where string, values ...interface{}) ([]*model.User, error) {

	users := make([]*model.User, 0)
	res := dao.MyDB.Table(tableUsers).Where(where, values...).Find(&users)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return users, nil
}

func Update(updateMap map[string]interface{}, where string, values ...interface{}) error {

	res := dao.MyDB.Table(tableUsers).Where(where, values...).Updates(updateMap)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
