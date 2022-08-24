package questionsCH

import (
	"MyStackoverflow/clickhouse"
	"fmt"
	"testing"
)

func TestAutoMigrate(t *testing.T) {

	clickhouse.Init()
	//err := Insert(Question{
	//	Qid:        1,
	//	Uid:        1,
	//	Title:      "Test",
	//	Body:       "Test",
	//	Time:       time.Now(),
	//	IsResolved: 0,
	//	Likes:      0,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	qs, err := List("qid > ?", 0)
	for _, q := range qs {
		fmt.Println(q, err)
	}
}
