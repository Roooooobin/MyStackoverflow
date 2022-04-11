package topicsdao

import (
	"MyStackoverflow/model"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	topic := model.Topic{
		TopicName: "Java",
		ParentId:  1,
	}
	if err := Insert(topic); err != nil {
		fmt.Println(err)
		return
	}
}
