package queue

import (
	"bytes"
	"encoding/json"
	"strconv"
	"github.com/streadway/amqp"
)

//配置
const (
	hostConfig = "127.0.0.1"
	portConfig = "5672"
	userConfig = "guest"
	PasswordConfig = "guest"	
)

//Queue 队列结构体
type Queue struct {

	Priority          string
	JobQueueName      string //工作队列名
	TopicExchangeName string //topic业务交换机
	//DeadExchangeName string //死信交换机

	config ConnConfig
	//连接
	conn    *amqp.Connection
	channel *amqp.Channel
	//错误
	err error
}

type ConnConfig struct {
	Host              string
	Port              string
	User              string
	Password          string
}

//队列对象
var MainQue Queue

//New 初始化默认
func New() (q *Queue) {
	q =  &Queue{
		Priority:          "8",
		JobQueueName:      "main",
		TopicExchangeName: "topic_exchange",
	}
	//配置
	q.SetConfig([]string{hostConfig,portConfig,userConfig,PasswordConfig})
	//创建MQ连接
	q.MakeConn()
	return
}


//NewConfig 初始化带配置
func NewConfig(config []string) (q *Queue) {
	q =  &Queue{
		Priority:          "8",
		JobQueueName:      "main",
		TopicExchangeName: "topic_exchange",
	}
	//配置
	q.SetConfig(config)
 	//创建MQ连接
 	q.MakeConn()
	return
}

//Destroy 关闭连接
func (q *Queue) Destroy() {
	q.channel.Close()
	q.conn.Close()
}

//MakeConn 创建连接
func (q *Queue) MakeConn() {
	q.conn, q.err = amqp.Dial("amqp://" + q.config.User + ":" + q.config.Password + "@" + q.config.Host + ":" + q.config.Port + "/")
	if q.err == nil {
		q.channel, q.err = q.conn.Channel()
	}
}


//SetConfig 自定义配置
func (q *Queue) SetConfig (config []string) *Queue {
	if len(config)>0 {
		q.config.Host = config[0]
		q.config.Port = config[1]
		q.config.User = config[2]
		q.config.Password = config[3]
	}
	return q
}

//SetQueue 选择队列，只能传在config里面配的 main task error
func (q *Queue) SetQueue(name string) *Queue {
	q.JobQueueName = name
	return q
}

//SetPriority 设置优先级 queue.New().SetPriority("6").Push("TestJob", map[string]interface{}{"Id": 1, "Name": "18103129"})
func (q *Queue) SetPriority(priority string) *Queue {
	q.Priority = priority
	return q
}


//Job 工作队列
type JobReceivers interface {
	Execute(interface{}) error //执行任务
}


//Push 工作队列模式 1、jobName工作名 2、data数据
func (q *Queue) Push(jobName string, data interface{}) error {

	if q.err != nil {
		return q.err
	}
	ch := q.channel

	queue, err := ch.QueueDeclare(
		q.JobQueueName, //name
		true,           //durable
		false,          //delete when unused
		false,          //exclusive
		false,          //no-wait
		nil,            //arguments
	)
	if err != nil {
		return err
	}
	queueBody := map[string]interface{}{"job": jobName, "data": data}
	body, _ := json.Marshal(queueBody)

	priority, err := strconv.Atoi(q.Priority)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",         //exchange
		queue.Name, //routing key
		false,      //mandatory
		false,      //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久化
			ContentType: "text/plain",
			Body:        []byte(body),
			Priority:    uint8(priority),
		})
	//fmt.Println("推送消息：", queueBody)
	//fmt.Println("队列名称：", q.JobQueueName)
	if err != nil {
		return err
	}
	q.Destroy()
	return nil
}

//Listen 工作队列监听 1、队列名queueName
func (q *Queue) Listen(Jobs map[string]JobReceivers) error {

	if q.err != nil {
		return q.err
	}
	ch := q.channel


	queue, err := ch.QueueDeclare(
		q.JobQueueName, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := bytesToString(&(d.Body))

			m := make(map[string]interface{})
			json.Unmarshal([]byte(*msg), &m)
			//fmt.Println(" 收到消息 ：", m)

			//断言job值为string，才可用于map的key
			if key, ok := m["job"].(string); ok {
				//判断是否存在该key名，才可执行Job的Execute，否则不存在的key会报错
				if oneJob, ok := Jobs[key]; ok {
					err := oneJob.Execute(m["data"])
					if err == nil {
						//fmt.Println("完成 ..")
					} else {
						//错误处理
						queue, errA := ch.QueueDeclare(
							"gw_error", //name
							true,       //durable
							false,      //delete when unused
							false,      //exclusive
							false,      //no-wait
							nil,        //arguments
						)
						if errA != nil {
							//fmt.Errorf("错误队列绑定失败", errA)
						}
						body, _ := json.Marshal(map[string]interface{}{"job": m["job"].(string), "data": m["data"] , "err":err.Error()})
						errB := ch.Publish(
							"",         //exchange
							queue.Name, //routing key
							false,      //mandatory
							false,      //immediate
							amqp.Publishing{
								DeliveryMode: amqp.Persistent, //持久化
								ContentType: "text/plain",
								Body:        []byte(body),
							})
						//fmt.Println("工作发生错误，错误队列推送消息：", m)
						if errB != nil {
							//fmt.Errorf("错误队列消息推送失败", errB)
						}
					}
					d.Ack(false)

				}
			}
			//d.Ack(true)
		}
	}()

	//fmt.Println(" 等待工作..")
	<-forever

	q.Destroy()
	return nil
}

//bytes转字符串
func bytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}


//topic接口
type TopicReceivers interface {
	GetQueueName() string
	GetRoutingKeys() []string
	Execute(routingKey string, data interface{}) error
}


//TopicPush  1、路由名 2、数据
func (q *Queue) TopicPush(routingKey string, data interface{}) error {

	if q.err != nil {
		return q.err
	}
	ch := q.channel


	err := ch.ExchangeDeclare(
		q.TopicExchangeName, // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return err
	}
	queueBody := map[string]interface{}{"key": routingKey, "data": data}

	body, _ := json.Marshal(queueBody)
	err = ch.Publish(
		q.TopicExchangeName, // exchange
		routingKey,          // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久化
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	//fmt.Println(" 发送消息 %s , 路由名 %s", body, routingKey)
	q.Destroy()
	return nil

}

//监听：queue.New().TopicListen(topic.Task{})
//TopicListen  队列监听 传入参数必须是topic类型的接口才可监听
func (q *Queue) TopicListen(t TopicReceivers) error {

	if q.err != nil {
		return q.err
	}
	ch := q.channel


	err := ch.ExchangeDeclare(
		q.TopicExchangeName, // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		t.GetQueueName(), // name
		true,             // durable
		false,            // delete when usused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return err
	}

	for _, s := range t.GetRoutingKeys() {
		////wlog.WithFields(logrus.Fields{"queue": queue.Name, "router": q.TopicExchangeName, "rule": s}).Info("绑定队列 %s 到交换机 %s ，路由规则为： %s", queue.Name, q.TopicExchangeName, s)
		err = ch.QueueBind(
			queue.Name,          // queue name
			s,                   // routing key
			q.TopicExchangeName, // exchange
			false,
			nil)
		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := bytesToString(&(d.Body))

			dataMap := make(map[string]interface{})
			json.Unmarshal([]byte(*msg), &dataMap)
			//fmt.Println(" 收到消息 %s", dataMap)

			err := t.Execute(d.RoutingKey, dataMap["data"])
			if err == nil {
				//fmt.Println("完成 ..")
			} else {

				//错误处理
				dataMap["err"] = err.Error()

				queue, err := ch.QueueDeclare(
					t.GetQueueName()+"_error", //name
					true,                      //durable
					false,                     //delete when unused
					false,                     //exclusive
					false,                     //no-wait
					nil,                       //arguments
				)
				if err != nil {
					//fmt.Errorf("错误队列绑定失败，", err)
				}

				body, _ := json.Marshal(dataMap)
				err = ch.Publish(
					"",         //exchange
					queue.Name, //routing key
					false,      //mandatory
					false,      //immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent, //持久化
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				if err != nil {
					//fmt.Errorf("错误队列消息推送失败", err)
				}

				//fmt.Println("工作发生错误,错误队列推送消息：", dataMap)
			}
			d.Ack(false)

		}
	}()

	//fmt.Println("等待工作...")
	<-forever

	q.Destroy()
	return nil
}


/* 
	var ServicesMap = make(map[string][]interface{})
	ServicesMap["wt_queue"] = []interface{}{"wt_queue",&service.AccountServer{},"SalSalary",&pb.AuthCheckRequest{},oneConfig}
	ServicesMap["oa_entry"] = []interface{}{"oa_entry",&service.IhrSalaryCheckServer{},"SalSalary",&pb.AuthCheckRequest{},oneConfig}
	tq := rb.Mq
	tq.MakeChannel()
	tq.Consumer(v[0].(string),v[1],v[2].(string),v[3],v[4].(rb.ReceiverConfig))
 */
 
 
/* 


//Pub 1、交换机 2、路由名 3、数据
func (q *Queue) Publish(exchangeName string,routingKey string, data interface{}) error {

	//如果没规划channel ,创建channel
	//q.MakeChannel()

	if q.err != nil {
		return q.err
	}
	ch := q.channel

	queueHeader := make(map[string]interface{},5)
	//queueHeader["retry"] = int32(0)
	queueHeader["routing"] = routingKey
	queueHeader["exchange"] = routingKey

	body, _ := json.Marshal(data)
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,          // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久化
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers:queueHeader,
		})
	if err != nil {
		return err
	}
	//
	log.Printf(" 发送消息 %s , 路由名 %s ，交换机  %s", body, routingKey,exchangeName)
	q.Destroy()
	return nil

}


//配置项
type ReceiverConfig struct {
	//流控
	PrefetchCount int
	PrefetchSize int
	Global bool
	//策略
	XCancelOnHaFailover bool

	//错误后是否ack
	ErrorAck bool
}


//ConsumerListenService  队列Service监听 执行  service服务的funcName方法
func (q *Queue) Consumer(queueName string, service interface{},funcName string, request interface{},rConfig ReceiverConfig) error {

	//如果没规划channel ,创建channel
	//q.MakeChannel()

	if q.err != nil {
		return q.err
	}
	ch := q.channel

	//解析接收者 配置
	ch.Qos(rConfig.PrefetchCount,0,false)
	var args = make(amqp.Table,10)
	args["x-cancel-on-ha-failover"] = rConfig.XCancelOnHaFailover

	msgs, err := ch.Consume(
		queueName, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		args,        // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := bytesToString(&(d.Body))
			log.Printf(" 收到消息 %s", *msg)

			//把消息解析进pb的request
			json.Unmarshal([]byte(*msg), request)

			//反射service服务并执行funcName方法，请求是ctx,request
			params:= make([]reflect.Value, 2)
			ctx := context.Background()
			params[0] = reflect.ValueOf(ctx)
			params[1] = reflect.ValueOf(request)
			out := reflect.ValueOf(service).MethodByName(funcName).Call(params)

			//如果执行return error，则记录错误
			if len(out)>1 {
				if out[1].Interface()!=nil {
					err = out[1].Interface().(error)
				}
			}

			//一条消息执行结束
			if err == nil {
				//成功
				d.Ack(false)
				log.Printf("完成 : %s",funcName)
			} else {
				//错误
				if rConfig.ErrorAck == true {
					d.Ack(false)
				}
				log.Printf("消费者发生错误 ： %s", err.Error())
			}

		}
	}()

	fmt.Println( "等待工作..." )
	<-forever

	q.Destroy()
	return nil
} */
