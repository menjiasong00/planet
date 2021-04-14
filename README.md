
## 原始依赖


获取依赖包、项目

```
[cmd]
    #获取依赖包
    $ go get -u github.com/philips/grpc-gateway-example
    
    #把项目clone到本地
    $ git clone github.com/menjiasong00/planet.git
```

 启动

```
[cmd]
    $go run main.go serve 
```



[gogrpc]: https://github.com/grpc/grpc-go
[grpcgateway]: https://github.com/grpc-ecosystem/grpc-gateway

# 一、planet 框架总览

1、基于 grpc-gateway ,同时满足http和rpc的api微服务框架
```
[.go]
    //微服务配置
    ServerSettings = []gcore.ServeSetting{
            //测试demo
            {
                Host:":"+grpcPort, //注册的微服务端口
                Server:&service.TestServer{}, //注册的微服务
                Register:pb.RegisterTestServer, 
                HandlerFromEndpoint:pb.RegisterTestHandlerFromEndpoint,
            },
            ...
        }
    gcore.RunServeGRPC(ServerSettings,grpcPort)
    gcore.RunServeHTTP(ServerSettings,httpPort)
```

2、基于 rabbitMq的微服务总线框架： 解耦各个微服务间的消息发布订阅交互  
```
[.go]
    //消费者配置
    ConsumerSettings = []grbmq.ConsumerSetting{
        //test
        {
            QueueName:"oa.employee.ihr",     //队列
            RoutingKey:"oa.employee.entry",  //逻辑路由
            Workers:2,                       //消费者数量
            Service:&service.TestServer{},   //微服务
            Controller:"GetTestMsg",         //消费方法
            Request:&pb.TestMessage{},       //消息体
            Config : grbmq.ReceiverConfig{1, 1, true, false, false, false},  //配置
        },
        ...
    }
    grbmq.Mq.RunConsumers(ConsumerSettings)
```

3、编写了一套api示例生成工具，快速开发一个业务对象的增删改查api

```
[.go]
    //代码
    gcode.MakeCoding(gcode.MakeCodingRequest{
		DatabaseName:"test",   //数据库名
		TableName:"products",  //数据表名
		Name:"产品",           //模块中文
		ServerName:"Bas",      //微服务名
		ModuleName:"BaslProducts", //模块编码
	})
```

# 二、api自动生成工具


1、安装mysql（新手自行百度，推荐使用phpstudy等工具快速安装）。

2、创建test测试库，创建products表，sql参考：
```
[.sql]
    //建表代码
    CREATE TABLE `products` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `created_at` timestamp NULL DEFAULT NULL,
      `updated_at` timestamp NULL DEFAULT NULL,
      `deleted_at` timestamp NULL DEFAULT NULL,
      `code` varchar(255) DEFAULT NULL,
      `price` int(10) unsigned DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `idx_products_deleted_at` (`deleted_at`)
    ) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
    
    INSERT INTO `test`.`products` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `price`) VALUES ('1', '2018-10-30 11:41:21', '2018-10-30 11:41:21', NULL, 'M3344', '2234');
    INSERT INTO `test`.`products` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `price`) VALUES ('2', '2018-10-30 11:43:38', '2018-10-30 11:43:38', NULL, 'L1212', '2234');

```

3、planet\service\testService.go 的注释去掉：
```
[.go]
    func (m *TestServer) GetTestMsg(c context.Context, s *pb.TestMessage) (*pb.TestMessage, error) {
        fmt.Printf("xxxxx(%q)\n", s.Value)
        //去掉这里的注释
        /*gcode.MakeCoding(gcode.MakeCodingRequest{
            DatabaseName:"test",
            TableName:"products",
            Name:"产品",
            ServerName:"Bas",
            ModuleName:"BaslProducts",
        })*/
        return s, nil
    }
```
且重启服务：

```
[cmd]
    $go run main.go serve 
```
4、调用测试api 。执行gcode.MakeCoding方法，生成对应的api代码
```
[url]
    http://127.0.0.1:8080/v1/test?value=testmycode
```
5、检查已生成的3个文件：并改文件名去掉 “-”
```
[file]
    planet\service\basService.go-
    planet\model\products.go-
    planet\proto\bas.proto-
```
6、执行脚本生成.proto的pb代码
```
[cmd]
    planet\proto\run.bat
```
7、将生成的微服务注册到项目中planet\cmd\serve.go：
```
[.go]
    ServerSettings = []gcore.ServeSetting{
            //测试demo
            {
                Host:":"+grpcPort,
                Server:&service.TestServer{},
                Register:pb.RegisterTestServer,
                HandlerFromEndpoint:pb.RegisterTestHandlerFromEndpoint,
            },
            //添加以下代码，注册 刚刚生成的bas微服务代码
            {
                Host:":"+grpcPort,
                Server:&service.BasServer{},
                Register:pb.RegisterBasServer,
                HandlerFromEndpoint:pb.RegisterBasHandlerFromEndpoint,
            },
        }
```


8、重启服务，查看接口
```
[url]
    http://127.0.0.1:8080/bas/products
```

# 三、微服务代码结构

根据上面的操作已经完成了一个api接口的生成。接下来讲解一个api的代码结构

1、Service层：逻辑代码：
```
[.go]
    //ProductsDetail 产品详情
    func (s *BasServer) ProductsDetail(ctx context.Context, in *pb.ProductsIdRequest) (*pb.ProductsDetailResponse, error) {
    
        resp := &pb.ProductsOneRequest{}
        gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Scan(&resp)
    
        return &pb.ProductsDetailResponse{Status: 200, Message: "success", Details: resp}, nil
    }
```

2、model层：数据库表结构
```
[.go]
    //产品
    type Products struct {
        Id int  `gorm:"column:id"`
        CreatedAt time.Time  `gorm:"column:created_at"`
        UpdatedAt time.Time  `gorm:"column:updated_at"`
        DeletedAt time.Time  `gorm:"column:deleted_at"`
        Code string  `gorm:"column:code"`
        Price int  `gorm:"column:price"`
    }
    
    //产品表名
    func (Products) TableName() string {
        return "products"
    }
```
3、proto层：api地址、请求和返回结构
```
[.go]
    //产品详情
    rpc ProductsDetail (ProductsIdRequest) returns (ProductsDetailResponse) {
        option (google.api.http) = {
            get: "/bas/products/{id}"
		};
    }
```

4、为了把proto自动编译go代码。 开发必备:安装go工具组件

```
[cmd]
    $go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    $go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    $go get -u github.com/golang/protobuf/protoc-gen-go
    $go get -u github.com/envoyproxy/protoc-gen-validate
    $go get -u github.com/golang/protobuf/proto // golang protobuf 库
    $go get google.golang.org/grpc
```
下载protobuf 3.6.1  

https://github.com/protocolbuffers/protobuf/releases

protoc 复制 到d:\gopath\bin google复制到planet\echopb 有写好的脚本run,把proto生成go的代码pb



# 四、消息框架的使用

1、案例说明：假设某微博平台： 用户微服务A 、积分微服务B 、通知推送微服务C。

2、业务场景：用户注册后，积分系统B 先为用户初始化积分，送300分、通知系统C 再通知给该用户可能认识的人。

3、生产端

生产者：用户微服务A收到用户注册表单，并向rabbimq发送一条消息，告诉其他B、C服务“用户已经注册了”：

```
    //业务上游 A 往交换机Exchange【 USER.PERSONS.PUBLISH 】发送了一条“注册成功”的消息，路由key为【 user.person.entry 】:
    grbmq.Mq.Publish("USER.PERSONS.PUBLISH", "user.person.entry", map[string]interface{}{"name": "test", "type": "entry", "msg": "注册成功"})
```

4、消费端

4.1、消费者B ：在rabbitMq建造一条队列为 【 user.person.goals 】，和 交换机Exchange【 USER.PERSONS.PUBLISH 】进行绑定，路由名为【 user.person.entry 】。并在 serve.go配置如下消费者：
```
    ConsumerSettings = []grbmq.ConsumerSetting{
        ...
        //积分系统消费者配置
        {
            QueueName:"user.person.goals",
            RoutingKey:"user.person.entry",
            Workers:2,
            Service:&service.GoalsServer{},       //积分微服务
            Controller:"PersonEntryGoalsConsumer",//消费方法
            Request:&pb.PersonEntryGoalsMessage{},//消息请求体
            Config : grbmq.ReceiverConfig{1, 1, true, false, false, false},
        },
        ...
    }
```

配置消费者后，在GoalsServer微服务中实现消费方法接口：
```
    func (m *GoalsServer) PersonEntryGoalsConsumer(c context.Context, s *pb.PersonEntryGoalsMessage) (*pb.TestMessage, error) {
        fmt.Printf("加积分逻辑\n", s.Value)
        //...
        return s, nil
    }
```

4.2、消费者C ：在rabbitMq建造一条队列为 【 user.person.notice 】，和 交换机Exchange【 USER.PERSONS.PUBLISH 】进行绑定，路由名为【 user.person.entry 】。并在 serve.go配置如下消费者：

```
    ConsumerSettings = []grbmq.ConsumerSetting{
        ...
        //通知系统消费者配置
        {
            QueueName:"user.person.notice",
            RoutingKey:"user.person.entry",
            Workers:2,
            Service:&service.NoticeServer{},       //通知微服务
            Controller:"PersonEntryNoticeConsumer",//消费方法
            Request:&pb.PersonEntryNoticeMessage{},//消息请求体
            Config : grbmq.ReceiverConfig{1, 1, true, false, false, false},
        },
        ...
    
    }
```
配置消费者后，在NoticeServer微服务中实现消费方法接口：
```
    func (m *NoticeServer) PersonEntryNoticeConsumer(c context.Context, s *pb.PersonEntryNoticeMessage) (*pb.TestMessage, error) {
        fmt.Printf("发通知逻辑\n", s.Value)
        //...
        return s, nil
    }
```


8、重启服务，发现2个队列已经加入监听。调用注册的接口，其他两个消费者就会执行打印：
```
加积分逻辑
发通知逻辑
```

# 五、其他


 打开go mod: $Env:GO111MODULE="on" 可以设置用户环境变量代理解决go get 慢问题:

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

设置host（C:\Windows\System32\drivers\etc）可以解决github无法访问的问题。编辑hosts：
```
# Github Hosts
# Update 20210312
140.82.113.4 github.com
140.82.112.10 nodeload.github.com
140.82.114.6 api.github.com
13.229.189.0 codeload.github.com
185.199.110.133 raw.github.com
185.199.110.153 training.github.com
185.199.110.153 assets-cdn.github.com
185.199.110.153 documentcloud.github.com
185.199.110.154 help.github.com

185.199.110.153 githubstatus.com
199.232.69.194 github.global.ssl.fastly.net

185.199.110.133 raw.githubusercontent.com
185.199.110.133 cloud.githubusercontent.com
185.199.110.133 gist.githubusercontent.com
185.199.110.133 marketplace-screenshots.githubusercontent.com
185.199.110.133 repository-images.githubusercontent.com
185.199.110.133 user-images.githubusercontent.com
185.199.110.133 desktop.githubusercontent.com

185.199.110.133 avatars.githubusercontent.com
185.199.110.133 avatars0.githubusercontent.com
185.199.110.133 avatars1.githubusercontent.com
185.199.110.133 avatars2.githubusercontent.com
185.199.110.133 avatars3.githubusercontent.com
185.199.110.133 avatars4.githubusercontent.com
185.199.110.133 avatars5.githubusercontent.com
185.199.110.133 avatars6.githubusercontent.com
185.199.110.133 avatars7.githubusercontent.com
185.199.110.133 avatars8.githubusercontent.com
185.199.108.153 assets-cdn.github.com
185.199.109.153 assets-cdn.github.com
185.199.110.153 assets-cdn.github.com
185.199.111.153 assets-cdn.github.com
# End of the section
```

 安装docker,启调用链监控jaeger

docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest

访问http://127.0.0.1:16686/

 计划整合:

Docker的图形化管理工具Portainer

Docker的数据仓库harbor

JWT

dockerSwarm

Exporter-系统监控Prometheus-可视化Grafana-Alerting