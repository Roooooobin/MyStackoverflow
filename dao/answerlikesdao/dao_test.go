package answerlikesdao

import (
	"MyStackoverflow/model"
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	like := model.AnswerLike{
		Uid: 1,
		Aid: 1,
	}
	if err := Insert(like); err != nil {
		fmt.Println(err)
		return
	}
}

func TestFind(t *testing.T) {

	like, err := Find("uid = ? and aid > ?", 1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(like.Uid, like.Aid, like.Time)

}

func TestList(t *testing.T) {

	likes, err := List("uid = ? and aid >= ?", 1, 2)
	if err != nil {
		fmt.Println(err)
	}
	for _, like := range likes {
		fmt.Println(like.Uid, like.Aid, like.Time)
	}
}

func TestUpdate(t *testing.T) {

	updateMap := map[string]interface{}{
		"aid":  4,
		"time": time.Now(),
	}
	err := Update(updateMap, "uid = ? and aid = ?", 1, 2)
	if err != nil {
		return
	}
}
