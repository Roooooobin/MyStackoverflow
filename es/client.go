package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func NewEsClient() *elastic.Client {

	url := fmt.Sprintf("http://%s:%d", "127.0.0.1", 9200)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	return client
}
