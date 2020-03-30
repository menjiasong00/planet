package queue

import (
	"bytes"
	"encoding/json"
	"planet/pkg/gqueue/topic"
	"planet/pkg/gqueue/jobs"
	"github.com/sirupsen/logrus"
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

func SetConfig (config map[string]string) {
	hostConfig = config[0]
	portConfig = config[1]
	userConfig = config[2]
	PasswordConfig = config[3]
}

var wlog *logrus.Logger

//Queue 队列结构体
type Queue struct {
	Host              string
	Port              string
	User              string
	Password          string
	Priority          string
	JobQueueName      string //工作队列名
	TopicExchangeName string //topic业务交换机
	//DeadExchangeName string //死信交换机

	//连接
	conn    *amqp.Connection
	channel *amqp.Channel
	//错误
	err error
}

//队列对象
var MainQue Queue

//New 初始化
func New() (q *Queue) {
	q =  &Queue{
		Host:              hostConfig,
		Port:              portConfig,
		User:              userConfig,
		Password:          PasswordConfig,
		Priority:          "8",
		JobQueueName:      "main"),
		TopicExchangeName: "topic_exchange",
	}

 	//创建MQ连接
	q.conn, q.err = amqp.Dial("amqp://" + q.User + ":" + q.Password + "@" + q.Host + ":" + q.Port + "/")
	if q.err == nil {
		q.channel, q.err = q.conn.Channel()
	}

	return
}

//Destroy 关闭连接
func (q *Queue) Destroy() {
	q.channel.Close()
	q.conn.Close()
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
	wlog.Info("推送消息：", queueBody)
	wlog.Info("队列名称：", q.JobQueueName)
	if err != nil {
		return err
	}
	q.Destroy()
	return nil
}

//Listen 工作队列监听 1、队列名queueName
func (q *Queue) Listen() error {

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
			wlog.Info(" 收到消息 ：", m)

			//断言job值为string，才可用于map的key
			if key, ok := m["job"].(string); ok {
				//判断是否存在该key名，才可执行Job的Execute，否则不存在的key会报错
				if oneJob, ok := Jobs[key]; ok {
					err := oneJob.Execute(m["data"])
					if err == nil {
						wlog.Info("完成 ..")
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
							wlog.Error("错误队列绑定失败", errA)
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
						wlog.Info("工作发生错误，错误队列推送消息：", m)
						if errB != nil {
							wlog.Error("错误队列消息推送失败", errB)
						}
					}
					d.Ack(false)

				}
			}
			//d.Ack(true)
		}
	}()

	wlog.Info(" 等待工作..")
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
	wlog.Info(" 发送消息 %s , 路由名 %s", body, routingKey)
	q.Destroy()
	return nil

}

//监听：queue.New().TopicListen(topic.Task{})
//TopicListen  队列监听 传入参数必须是topic类型的接口才可监听
func (q *Queue) TopicListen(t topic.Topic) error {

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
		//wlog.WithFields(logrus.Fields{"queue": queue.Name, "router": q.TopicExchangeName, "rule": s}).Info("绑定队列 %s 到交换机 %s ，路由规则为： %s", queue.Name, q.TopicExchangeName, s)
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
			wlog.Info(" 收到消息 %s", dataMap)

			err := t.Execute(d.RoutingKey, dataMap["data"])
			if err == nil {
				wlog.Info("完成 ..")
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
					wlog.Error("错误队列绑定失败，", err)
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
					wlog.Error("错误队列消息推送失败", err)
				}

				wlog.Info("工作发生错误,错误队列推送消息：", dataMap)
			}
			d.Ack(false)

		}
	}()

	wlog.Info("等待工作...")
	<-forever

	q.Destroy()
	return nil
}
