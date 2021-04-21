package cmd

import (
	"planet/env"
	"planet/env/server"
	"planet/pkg/gcore"
	//"planet/insecure"
	"github.com/spf13/cobra"
)

func init() {
	//服务
	RootCmd.AddCommand(serveCmd)
	//网关
	RootCmd.AddCommand(gatewayCmd)
	//单体服务
	RootCmd.AddCommand(DevCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the serve ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)>0 {
			servFlag := args[0]
			grpcPort := env.Config.GetString("Server."+servFlag+".GrpcPort")
			//各个rpc服务
			gcore.RunServeGRPC(server.ConfigServer(servFlag),grpcPort)
			/*var rbmqConfig grbmq.ConnConfig
			env.ScanConfig("Rbmq",&rbmqConfig)
			grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/
		}
	},
}

// serveCmd represents the serve command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Launches the gateway  ",
	Run: func(cmd *cobra.Command, args []string) {
		servFlag := "GateWay"
		httpPort := env.Config.GetString("Server."+servFlag+".HttpPort")
		gcore.RunServeHTTP(server.ConfigServer(servFlag),httpPort)
	},
}

// serveCmd represents the dev command
var DevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Launches the example webserver  ",
	Run: func(cmd *cobra.Command, args []string) {
		//开发环境直接起
		servFlag := "Dev"
		grpcPort := env.Config.GetString("Server.Dev.GrpcPort")
		httpPort := env.Config.GetString("Server.Dev.HttpPort")

		/*var rbmqConfig grbmq.ConnConfig
		env.ScanConfig("Rbmq",&rbmqConfig)
		grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/
		//注册所有
		go gcore.RunServeGRPC(server.ConfigServer(servFlag),grpcPort)
		gcore.RunServeHTTP(server.ConfigServer(servFlag),httpPort)
		//gcore.MakeInsecure(insecure.Key,insecure.Cert)
		//serve()
	},
}