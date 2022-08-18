package clickhouse

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var clickhouseDB *gorm.DB

func Init() {

	dsn := "tcp://localhost:9000?database=MyStackoverflow&username=default&password="
	var err error
	clickhouseDB, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
