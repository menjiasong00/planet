
#1 planet是一个基于 grpc-gateway ,同时满足http和rpc的框架,集成了 微服务\接口文档\日志\认证\调用链等基础服务.你可以在这个基础上快速开发使用微服务+http json接口

git clone github.com/menjiasong00/planet

#2 打开go mod: $Env:GO111MODULE="on" 可以设置用户环境变量代理解决go get 慢问题:

GOPROXY=https://goproxy.io

GO111MODULE=on

#4 启动

go run main.go serve 10000

访问:https://localhost:10000/v1/test?value=44444
接口文档:https://localhost:10000/swagger-ui/

#5 开发必备:安装工具组件

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/envoyproxy/protoc-gen-validate

#6 下载protobuf 3.6.1 修改proto后 运行工具生成go代码

https://github.com/protocolbuffers/protobuf/releases 

protoc 复制 到d:\gopath\bin  google复制到planet\echopb 有写好的脚本run生成pb

#6 安装docker,启调用链监控jaeger 

docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest

访问http://127.0.0.1:16686/  



#计划整合:

Docker的图形化管理工具Portainer

Docker的数据仓库harbor

JWT

dockerSwarm

Exporter-系统监控Prometheus-可视化Grafana-Alerting



