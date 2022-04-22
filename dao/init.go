package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyDB = getDB()

func getDB() *gorm.DB {

	username := "root"
	password := "19980506"
	host := "127.0.0.1"
	port := 3306
	dbName := "my_stackoverflow"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, dbName, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failed, error=" + err.Error())
	}
	return db
}
