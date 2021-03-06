package main

import (
	"fmt"
	que "github.com/menjiasong00/queue"
)

func main() {

	oneTopicQue := que.NewConfig([]string{"10.10.18.130","5672","guest","guest"})

	oneTopicQue.TopicQueueBind("topic_test",[]string{"xx.*","xx.22.xx"})

	err :=oneTopicQue.TopicListen(MsgTopic{})
	
	fmt.Println(err.Error())

}

type MsgTopic struct {}

// 执行发邮件
func (c MsgTopic) GetQueueName() string {
	return "topic_test"
}

// 执行发邮件
func (c MsgTopic) Execute(routingKey string,data interface{}) error {
	fmt.Println(routingKey)
	fmt.Println(data)
	return nil
}
