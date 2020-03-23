protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis    --go_out=plugins=grpc:./../pb   service.proto
protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --grpc-gateway_out=logtostderr=true:./../pb service.proto
protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --swagger_out=logtostderr=true:./../doc service.proto


protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis    --go_out=plugins=grpc:./../pb   test.proto
protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --grpc-gateway_out=logtostderr=true:./../pb test.proto
protoc -I. -IC:\gopath\src\protobuf  -IC:\gopath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --swagger_out=logtostderr=true:./../doc test.proto