package gcore

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"reflect"
)

func   ProtocValidateInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		out := reflect.ValueOf(req).MethodByName("Validate").Call(make([]reflect.Value, 0))
		if len(out)>0 {
			if out[0].Interface()!=nil {
				err = out[0].Interface().(error)
			}
		}
		if err!=nil {
			return nil,err
		}
		resp, err = handler(ctx, req)
		return
	}
}