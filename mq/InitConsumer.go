package mq

import (
	"MyStackoverflow/dao/userstatsdao"
	"MyStackoverflow/model"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var C rocketmq.PushConsumer

func formProblemStats(uid int, body string) error {

	length := len(body)
	spaceSep := regexp.MustCompile("\\s+")
	words := spaceSep.Split(body, -1)
	wordsCount := len(words)
	userStats := model.UserStats{
		Uid:             uid,
		QuestionsLength: length,
		QuestionsWords:  wordsCount,
		AnswersLength:   0,
		AnswersWords:    0,
	}
	err := userstatsdao.InsertOrUpdate(userStats)
	return err
}

func InitConsumer() {

	C, _ = rocketmq.NewPushConsumer(
		// 消费组
		consumer.WithGroupName("stackoverflow"),
		// nameserver地址
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	// subscribe before consume
	err := C.Subscribe("questions", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		// consume problem and store statistics
		for i := range ext {
			fmt.Printf("%s\n", ext[i].Body)
			msg := string(ext[i].Body)
			ss := strings.Split(msg, ":")
			uidStr, body := ss[0], ss[1]
			uid, _ := strconv.Atoi(uidStr)
			errConsume := formProblemStats(uid, body)
			if errConsume != nil {
				return consumer.Rollback, errConsume
			}
		}
		//fmt.Println("consumeeeeeeeeeee")
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Printf("subscribe error: %s", err.Error())
	}
	err = C.Start()
	if err != nil {
		fmt.Printf("start error: %s", err.Error())
		os.Exit(-1)
	}
}
