package cmd

import (
	"planet/env"
	"planet/env/server"
	"planet/pkg/gcore"
	//"planet/insecure"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(DevCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on  "+demoAddr,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)>0 {
			servFlag := args[0]
			grpcPort := env.Config.GetString("Server."+servFlag+".GrpcPort")
			httpPort := env.Config.GetString("Server."+servFlag+".HttpPort")
			if servFlag == "GateWay" {
				//网关
				gcore.RunServeHTTP(server.ConfigServer(servFlag),httpPort)
			}else{
				//各个服务
				/*var rbmqConfig grbmq.ConnConfig
				env.ScanConfig("Rbmq",&rbmqConfig)
				grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/
				gcore.RunServeGRPC(server.ConfigServer(servFlag),grpcPort)

			}
		}





	},
}


// serveCmd represents the serve command
var DevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Launches the example webserver on  "+demoAddr,
	Run: func(cmd *cobra.Command, args []string) {
		//开发环境直接起
		servFlag := "Dev"
		grpcPort := env.Config.GetString("Server.Dev.GrpcPort")
		httpPort := env.Config.GetString("Server.Dev.HttpPort")

		/*var rbmqConfig grbmq.ConnConfig
		env.ScanConfig("Rbmq",&rbmqConfig)
		grbmq.New(rbmqConfig).RunConsumers(mq.ConfigConsumer)*/

		go gcore.RunServeGRPC(server.ConfigServer(servFlag),grpcPort)
		gcore.RunServeHTTP(server.ConfigServer(servFlag),httpPort)
		//gcore.MakeInsecure(insecure.Key,insecure.Cert)
		//serve()

	},
}