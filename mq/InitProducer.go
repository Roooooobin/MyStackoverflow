package mq

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var P rocketmq.Producer

func InitProducer() {

	P, _ = rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
		producer.WithGroupName("stackoverflow"),
	)
	err := P.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		panic(err)
	} else {
		fmt.Println("success")
	}
	//topic := "problems"
	//for i := 0; i < 10; i++ {
	//	msg := &primitive.Message{
	//		Topic: topic,
	//		Body:  []byte("Hello RocketMQ Go Client" + strconv.Itoa(i)),
	//	}
	//	// 发送信息
	//	res, err := p.SendSync(context.Background(), msg)
	//	if err != nil {
	//		fmt.Printf("send message error:%s\n", err)
	//	} else {
	//		fmt.Printf("send message success: result=%s\n", res.String())
	//	}
	//}
	//err = p.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown producer error:%s", err.Error())
	//}

	//time.Sleep(time.Hour)
	//err = c.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown Consumer error:%s", err.Error())
	//}
}
