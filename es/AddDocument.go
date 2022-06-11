package main

import (
	"MyStackoverflow/model"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func AddDocument(client *elastic.Client) {
	user1 := model.User{
		Uid:      1,
		Username: "Robin",
		Status:   "Basic",
		Email:    "Robin23@gmail.com",
		Password: "123456",
		City:     "NY",
		State:    "NY",
		Country:  "US",
		Profile:  "Nothing",
	}
	ctx := context.Background()
	put1, err := client.Index().Index("users").Id("1").BodyJson(user1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(put1.Index, put1.Id)
}
