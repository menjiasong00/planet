package cmd

import (
	"planet/env"
	"planet/pb"
	"planet/pkg/gcore"
	"planet/pkg/grbmq"
	"planet/service"
	//"planet/insecure"
	"github.com/spf13/cobra"
)



var ConsumerSettings  = []grbmq.ConsumerSetting{}
var ServerSettings = []gcore.ServeSetting{}
var grpcPort = env.Config.GetString("Server.GrpcPort")
var httpPort = env.Config.GetString("Server.HttpPort")


func init() {
	//微服务配置
	ServerSettings = []gcore.ServeSetting{
		//测试demo
		{
			Host:":"+grpcPort,
			Server:&service.TestServer{},
			Register:pb.RegisterTestServer,
			HandlerFromEndpoint:pb.RegisterTestHandlerFromEndpoint,
		},
/*		{
			Host:":"+grpcPort,
			Server:&service.BasServer{},
			Register:pb.RegisterBasServer,
			HandlerFromEndpoint:pb.RegisterBasHandlerFromEndpoint,
		},*/
		{
			Host:":"+grpcPort,
			Server:&service.WebServer{},
			Register:pb.RegisterWebServer,
			HandlerFromEndpoint:pb.RegisterWebHandlerFromEndpoint,
		},
	}

	//消费者配置
	ConsumerSettings = []grbmq.ConsumerSetting{
	    //死信消费者
		/*{
			QueueName:"dlx.global.queue",
			RoutingKey:"dlx.global.queue",
			Workers:2,
			Service:&service.TestServer{},
			Controller:"DlxConsumer",
			Request:&pb.DlxConsumerRequest{},
			Config : rb.ReceiverConfig{1, 1, true, false, false, true},
		},*/
		//ihr
		{
			QueueName:"oa.employee.ihr",
			RoutingKey:"oa.employee.entry",
			Workers:1,
			Service:&service.TestServer{},
			Controller:"GetTestMsg",
			Request:&pb.TestMessage{},
			Config : grbmq.ReceiverConfig{1, 1, true, false, false, false},
		},

	}

	RootCmd.AddCommand(serveCmd)
}


// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on  "+demoAddr,
	Run: func(cmd *cobra.Command, args []string) {

		/*var rbmqConfig grbmq.ConnConfig
		env.ScanConfig("Rbmq",&rbmqConfig)
		grbmq.New(rbmqConfig).RunConsumers(ConsumerSettings)*/

		go gcore.RunServeGRPC(ServerSettings,grpcPort)
		gcore.RunServeHTTP(ServerSettings,httpPort)
		//gcore.MakeInsecure(insecure.Key,insecure.Cert)
		//serve()

	},
}