
1 go get github.com/menjiasong00/planet

2打开go mod: $Env:GO111MODULE="on" 可以设置用户环境变量代理解决go get 慢问题:

GOPROXY=https://goproxy.io
GO111MODULE=on


3 下载proto3.6.1
https://github.com/protocolbuffers/protobuf/releases 
protoc 复制 到d:\gopath\bin  google复制到planet\echopb
echopb有脚本生成pb

4 
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/envoyproxy/protoc-gen-validate

go run main.go serve 10000


调用连
docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest