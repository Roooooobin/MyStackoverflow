package es

import (
	"MyStackoverflow/model"
	"context"
	"fmt"
	"strconv"
)

func AddQuestion(question *model.Question) {

	ctx := context.Background()
	put1, err := esClient.Index().Index("question").Id(strconv.Itoa(question.Qid)).BodyJson(question).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(put1.Index, put1.Id)
}
