基于rabbitmq的 消息队列封装 

一、安装使用 
1、go get github.com/menjiasong00/queue (或者 git clone github.com/menjiasong00/queue.git)

2、测试

cd  github.com/menjiasong00/queue/test_job_push  

go run main.go 

二、两种模式试用 
1、工作模式 jobs

新起一个控制台

生产者:que.NewConfig([]string{"127.0.0.1","5672","guest","guest"}).Push("TestJob","xxxxxx")

代码例子：

cd /你的目录/github.com/menjiasong00/queue/test_job_push  

go run main.go 

消费者：que.NewConfig([]string{"127.0.0.1","5672","guest","guest"}).Listen(map[string]que.JobReceivers{"TestJob":MsgJob{}})

代码例子：

cd /你的目录/github.com/menjiasong00/queue/test_job_listen  

go run main.go 

2、主题订阅 topic 例子同上

三、代码说明
1、工作模式 jobs
我们看工作的接口

//Job 工作队列
type JobReceivers interface {
	Execute(interface{}) error //执行任务
}


推送发邮件的消息：

Push("SendEmail","this is an email")

在消息Push进队列后 ，监听程序Listen到消息，解析出map里Job名称 ，并调用对应的Execute

因此，只需要把工作的 接口实现出来：

type SendEmailJob struct {}

func (c SendEmailJob) Execute(data interface{}) error {
	// 业务代码
	fmt.Println(data)
	return nil
}

并在运行的进程Listen监听他  Listen(map[string]que.JobReceivers{"SendEmail":SendEmailJob{}}) 

把示例的生产者和消费者(可参考 queue/test_job_push 和queue/test_job_listen )go run main.go 。可以看到消费者中执行了SendEmailJob的 Execute

2、topic 主题订阅

我们看接受主题的接口

type TopicReceivers interface {
	GetQueueName() string
	GetRoutingKeys() []string
	Execute(routingKey string, data interface{}) error
}

推送邮件已写完的消息：

TopicPush("emain.write.finish","i finish an email")

在消息Push进队列后 ，N个接收者以 某些规则订阅了该消息，则这些接收者都能收到该消息： 

因此，只需要把接收者 接口实现出来：

接收者1：

type MsgTopic struct {}

func (c MsgTopic) GetQueueName() string {
	return "topic_email"
}

// 路由规则
func (c MsgTopic) GetRoutingKeys() []string {
	return []string{"emain.write.finish","emain.write.*"}
}

// 执行
func (c MsgTopic) Execute(routingKey string,data interface{}) error {
	fmt.Println(routingKey)
	fmt.Println(data)
	return nil
}

接收者2：

type TodoTopic struct {}

func (c TodoTopic) GetQueueName() string {
	return "topic_email"
}

// 路由规则
func (c TodoTopic) GetRoutingKeys() []string {
	return []string{"emain.write.finish","emain.write.*"}
}

// 执行
func (c TodoTopic) Execute(routingKey string,data interface{}) error {
	fmt.Println(routingKey)
	fmt.Println(data)
	return nil
}

并在运行的进程运行它们(可参考 queue/test_topic_push 和queue/test_topic_listen ) 

 TopicListen(MsgTopic{})  

 TopicListen(TodoTopic{}) 

两个进程都收到了消息并执行了对应的业务 Execute

三、总结

不管是工作模式还是订阅模式 ，设计思路都是 留出接口 让业务代码进行水平扩展，这样在业务中只需要去实现一个个 job或者topic的接收者。而不用关心消息的流转过程和处理。





