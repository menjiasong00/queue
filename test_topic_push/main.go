package main

import (

	que "github.com/menjiasong00/queue"
)

func main() {

	que.NewConfig([]string{"127.0.0.1","5672","guest","guest"}).TopicPush("xx.22.xx","test2222")

}