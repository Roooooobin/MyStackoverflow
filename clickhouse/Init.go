package clickhouse

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var ClickHouseDB *gorm.DB

func Init() {

	dsn := "tcp://localhost:9000?database=my_stackoverflow&username=default&password=&read_timeout=10s&write_timeout=20s"
	var err error
	ClickHouseDB, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
