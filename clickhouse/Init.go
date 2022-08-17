package main

import (
	"database/sql"
	"github.com/ClickHouse/clickhouse-go/v2"
	"time"
)

var clickhouseClient *sql.DB

func Init() {

	clickhouseClient = clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
	})
	clickhouseClient.SetMaxIdleConns(5)
	clickhouseClient.SetMaxOpenConns(10)
	clickhouseClient.SetConnMaxLifetime(time.Hour)
}

func main() {
	Init()
	//res, err := clickhouseClient.Exec("create table test (col1 int, col2 varchar(100)) engine = Memory")
	//if err != nil {
	//	return
	//}
	//fmt.Println(res)
	clickhouseClient.Exec("INSERT INTO test (col1, col2) values (1, '111')")
}
