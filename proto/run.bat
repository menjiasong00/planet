protoc -I. -IC:\Users\Administrator\go\src\protobuf  -IC:\Users\Administrator\go\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  -IC:/Users/Administrator/go/src/github.com/envoyproxy/protoc-gen-validate --go_out=plugins=grpc:./../pb --validate_out=lang=go:./../pb test.proto
protoc -I. -IC:\Users\Administrator\go\src\protobuf  -IC:\Users\Administrator\go\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --grpc-gateway_out=logtostderr=true:./../pb test.proto
protoc -I. -IC:\Users\Administrator\go\src\protobuf  -IC:\Users\Administrator\go\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --swagger_out=logtostderr=true:./../doc test.proto



::protoc -I. -ID:\GoPath\pkg\protoc-3.6.1-win32\include\google\protobuf  -ID:\GoPath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  -ID:/GoPath/src/github.com/envoyproxy/protoc-gen-validate --go_out=plugins=grpc:./../pb --validate_out=lang=go:./../pb *.proto
::protoc -I. -ID:\GoPath\pkg\protoc-3.6.1-win32\include\google\protobuf  -ID:\GoPath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --grpc-gateway_out=logtostderr=true:./../pb *.proto
::protoc -I. -ID:\GoPath\pkg\protoc-3.6.1-win32\include\google\protobuf  -ID:\GoPath\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis  --swagger_out=logtostderr=true:./../doc *.proto