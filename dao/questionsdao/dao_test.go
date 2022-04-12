package questionsdao

import (
	"MyStackoverflow/model"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	question := model.Question{
		Uid:   2,
		Title: "How to learn Java",
		Body:  "I want to learn Java, can anyone give me some advice?",
	}
	if err := Insert(question); err != nil {
		fmt.Println(err)
		return
	}
}
