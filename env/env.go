package env

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {

	//读取yaml文件
	Config = viper.New()
	//设置读取的配置文件
	Config.SetConfigName("config")
	//添加读取的配置文件路径
	Config.AddConfigPath("./env/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	Config.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig();err != nil {

	}

}