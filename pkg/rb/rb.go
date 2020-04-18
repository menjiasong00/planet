package rb

import (
	"bytes"
	"encoding/json"
	"github.com/rs/xid"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"reflect"
	"rest/config"
	"rest/util/tools"
	"rest/util/wlog"
	"time"
)


//Rbmq 结构体
type Rbmq struct {
	MsgHeader map[string]interface{}
	Priority  string

	config ConnConfig
	//连接
	conn    *amqp.Connection
	channel *amqp.Channel
	//错误
	err error

	//断线监控,重连
	notifyClose   chan *amqp.Error
	isConnected   bool
	done          chan bool

}

type ConnConfig struct {
	Host              string
	Port              string
	User              string
	Password          string
}

//Mq 全局静态mq变量
var Mq *Rbmq


func init() {
	if Mq == nil {
		Mq  = New().MakeConn()

	}
}

//New 初始化默认
func New() (q *Rbmq) {
	qConfig := ConnConfig{
		Host:config.Vip.GetString("queue.host"),
		Port:config.Vip.GetString("queue.port"),
		User:config.Vip.GetString("queue.username"),
		Password:config.Vip.GetString("queue.password"),
	}
	q =  &Rbmq{config:qConfig}
	q.Priority = "8"
	q.MsgHeader = make(map[string]interface{},10)
	appid:=config.Vip.GetString("queue.appid."+ *config.Server )
	if len(appid)>0 {
		q.MsgHeader["appid"] = config.Vip.GetString("queue.appid."+ *config.Server )
	}
	return
}

//ConnWatch 共享连接守护 官网建议,多线程尽量共享连接
func (q *Rbmq) connWatch(addr string) {
	defer func() {
		//捕获抛出的panic 进入重试， 不要影响进程挂掉
		if err := recover();err!=nil{
			wlog.Error("发生panic错误：",err)
			//log.Println(err)
		}
	}()

	for {
		q.isConnected = false
		for q.tryConnect(addr) != nil {
			err := q.tryConnect(addr)
			//log.Printf("connect false . retry... ： %s", err.Error())
			x:=config.Vip.GetInt("queue.retryWaitSecond")
			tryWaitSecond := time.Duration(x) * time.Second
			time.Sleep(tryWaitSecond)
			wlog.Error("mq连接错误，重试..",err)
		}
		select {
		case <-q.done:
			return
		case <-q.notifyClose:
		}
	}
}

//连接
func (q *Rbmq) tryConnect(addr string) error {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q.conn = conn
	q.channel = ch

	// 注册监控
	q.notifyClose = make(chan *amqp.Error)
	ch.NotifyClose(q.notifyClose)

	q.isConnected = true
	wlog.Info("mq已连接")
	//log.Println("mq has connect.")
	return nil
}

//bytes转字符串
func bytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

//SetConfig 自定义连接配置
func (q *Rbmq) SetConfig (config []string) *Rbmq {
	if len(config)>0 {
		q.config.Host = config[0]
		q.config.Port = config[1]
		q.config.User = config[2]
		q.config.Password = config[3]
	}
	return q
}

func (q *Rbmq) SetHeader (header  map[string]interface{}) *Rbmq {
	q.MsgHeader = header
	return q
}

//MakeConn 创建连接
func (q *Rbmq) MakeConn() *Rbmq{
	q.conn, q.err = amqp.Dial("amqp://" + q.config.User + ":" + q.config.Password + "@" + q.config.Host + ":" + q.config.Port + "/")
	return q
}

//MakeConn 创建channel
/*func (q *Rbmq) MakeChannel() *Rbmq{
	if q.conn !=nil{
		q.channel, q.err = q.conn.Channel()
	}
	return q
}*/

//Destroy 关闭
func (q *Rbmq) Destroy() {
	q.channel.Close()
	q.conn.Close()
}

//SetPriority 设置优先级
func (q *Rbmq) SetPriority(priority string) *Rbmq {
	q.Priority = priority
	return q
}


//MakeConn 获取一个channel
func (q *Rbmq) GetChannel()  (*amqp.Channel,error){
	return q.conn.Channel()
}


//Pub 1、交换机 2、路由名 3、数据
func (q *Rbmq) Publish(exchangeName string,routingKey string, data interface{}) error {
	q.MakeConn()
/*	//如果没有连接 创建连接

	//如果没规划channel ,创建channel
	q.MakeChannel()*/

	//官网建议,多线程尽量共享连接，独享信道。
	var ch *amqp.Channel
	if q.conn != nil {
		ch, q.err = q.conn.Channel()
	}
	if q.err != nil {
		return q.err
	}
	defer ch.Close()

	q.MsgHeader["key"] =  xid.New().String()
	q.MsgHeader["exchange"] = exchangeName
	q.MsgHeader["routing"] = routingKey

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
			Headers:q.MsgHeader,
		})
	if err != nil {
		return err
	}
	//
	wlog.Info("发送消息：",data,routingKey,exchangeName)
	//log.Printf(" send mgs: %s , routing: %s ，exchange:  %s", body, routingKey,exchangeName)
	//q.Destroy()
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
	//是否包含头
	IncludeHeader bool
}

//RunConsumer 起服务时候用
func (q *Rbmq)  RunConsumer(ServicesMap map[string][]interface{}) {

	//连接守护
	go Mq.connWatch("amqp://" + Mq.config.User + ":" + Mq.config.Password + "@" + Mq.config.Host + ":" + Mq.config.Port + "/")

 	for _,v:=  range  ServicesMap  {
		go q.Consumer(v[0].(string),v[1],v[2].(string),v[3],v[4].(ReceiverConfig))
	}

	time.Sleep(1*time.Second)
}


//Consumer  队列Service监听 执行  service服务的funcName方法
func (q *Rbmq) Consumer(queueName string, service interface{},funcName string, request interface{},rConfig ReceiverConfig) {

	//官网建议,多线程尽量共享连接，独享信道。
	var ch *amqp.Channel
	if q.conn != nil {
		ch, q.err = q.conn.Channel()
	}

	if q.conn != nil  && ch!=nil  && q.err ==nil {
		forever := make(chan bool)
		//解析接收者 配置
		ch.Qos(rConfig.PrefetchCount, 0, false)

		var args= make(amqp.Table, 10)
		args["x-cancel-on-ha-failover"] = rConfig.XCancelOnHaFailover

		msgs, err := ch.Consume(
			queueName, // queue
			"",        // consumer
			false,     // auto ack
			false,     // exclusive
			false,     // no local
			false,     // no wait
			args,      // args
		)
		if err != nil {
			wlog.Error("Consume错误：",err)
			//log.Printf("consume return error ： %s", err.Error())
			q.err = err
		}else{

		//异常监控,已经移动到公共的连接守护
		/*		go func(chan<- bool) {
					cc := make(chan *amqp.Error)
					err := <-q.channel.NotifyClose(cc)
					if err != nil {
						fmt.Printf("NotifyClose获取消费通道异常:%s \n", err)
						forever <- false
						return
					}
					forever <- true
				}(forever)*/

			go func(<-chan bool) {

				for {
					select {
					case d,chStatus := <-msgs:
						if !chStatus {
							wlog.Error("连接错误或者队列被更改：",err)
							//log.Printf("chanel error ,connection error or queue has been changed  :%s \n", d)
							forever <- false
							return
						}
/*					case d := <-msgs:
						//fmt.Println(d)
						if d.Headers == nil && d.Body == nil {
							log.Printf("empty msg ，discard :%s \n", d)
							err = d.Ack(true)
							forever <- false
							return
						}*/

						defer func() {
							//捕获抛出的panic 进入重试， 不要影响进程挂掉
							if err := recover();err!=nil{
								tools.FatalAndEmail("Mq消费发生panic的致命错误,消费者挂了："+queueName+" "+funcName,err)
								d.Reject(false)
								forever <- false
							}
						}()

						//消费
						msg := bytesToString(&(d.Body))
						wlog.Info("队列收到消息：",*msg,d.Headers,queueName,funcName)
						//log.Printf(" consumer queue: %s receive msg: %s", queueName, *msg)

						requestData:=[]byte(*msg)

						//请求体是否包含头
						if rConfig.IncludeHeader {
							headerJson, _ := json.Marshal(d.Headers)
							header := bytesToString(&(headerJson))
							requestData,_= json.Marshal(map[string]interface{}{"header":*header,"body":*msg})
						}

						//拷贝一个新的 指针变量
						requestNew := reflect.New(reflect.ValueOf(request).Elem().Type()).Interface()


						//把消息解析进pb的request
						json.Unmarshal(requestData, requestNew)

						//反射service服务并执行funcName方法，请求是ctx,request
						params := make([]reflect.Value, 2)
						ctx := context.Background()
						params[0] = reflect.ValueOf(ctx)
						params[1] = reflect.ValueOf(requestNew)
						out := reflect.ValueOf(service).MethodByName(funcName).Call(params)
						var errReturn error
						 //如果执行return error，则记录错误
						if len(out) > 1 {
							if out[1].Interface() != nil {
								errReturn = out[1].Interface().(error)
							}
						}
						//消费结束
						if errReturn != nil {
							//消费失败
							wlog.Error("方法执行错误：",errReturn, queueName, funcName,*msg , d.Headers)
							//log.Printf(" consumer queue: %s call func: %s , return error :%s , params: %s , header: %s\n", queueName, funcName, err,*msg , d.Headers)
							//手工拒绝 进入死信
							d.Reject(false)
						} else {
							errReturn = d.Ack(true)
							if errReturn != nil {
								//ack失败
								//wlog.Error("ack错误：",err, queueName, funcName,*msg , d.Headers)
								//log.Printf("ack false:%s %s \n", err,queueName)
								tools.FatalAndEmail("Mq消费ack错误,消费者挂了："+queueName+" "+funcName,errReturn)
								//forever <- false
								return
							}
							//消费成功
							wlog.Info("方法执行成功：",*msg,queueName,funcName)
							//log.Printf(" consumer queue: %s call func: %s , done , params: %s . ", queueName, funcName,*msg)
						}

					case <-q.notifyClose://守护连接协程发出异常
						forever <- false
						return
					case <-forever://队列内部异常
						return
					}
				}
			}(forever)

			wlog.Info("队列已启动，正在监听消息：",queueName,funcName)
			//log.Printf("consumer queue: %s waiting for msg ,listen func : %s", queueName,funcName)
			<-forever
			q.Destroy()
		}
	}

	wlog.Info("队列启动重试中...：",queueName,funcName)
	//log.Printf("consumer queue: %s will connect latter...", queueName)

	tryWaitSecond := time.Duration(config.Vip.GetInt("queue.retryWaitSecond")) * time.Second
	time.Sleep(tryWaitSecond)
	go q.Consumer(queueName , service ,funcName , request ,rConfig)
	//q.Destroy()
}
