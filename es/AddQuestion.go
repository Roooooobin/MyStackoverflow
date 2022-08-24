package es

import (
	"MyStackoverflow/model"
	"context"
	"strconv"
)

func AddQuestion(question *model.Question) error {

	ctx := context.Background()
	_, err := esClient.Index().Index("question").Id(strconv.Itoa(question.Qid)).BodyJson(question).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
