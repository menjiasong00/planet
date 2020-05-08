package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	//"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	//"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"planet/config/srvs"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"planet/pkg/ui/data/swagger"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	//"github.com/g4zhuj/grpc-wrapper/plugins"

	pb "planet/pb"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on https://localhost:10000",
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

type myService struct{}

func (m *myService) Echo(c context.Context, s *pb.EchoMessage) (*pb.EchoMessage, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Value)
	return s, nil
}

func newServer() *myService {
	return new(myService)
}


func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serve(args []string) {
	ServerMap := srvs.ServerMap
	portHost :=10000
	if len(args)>0 {
		portHost ,_ = strconv.Atoi(args[0])
	}
	
	//启动GRPC服务
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:"+strconv.Itoa(int(portHost))))}
 
 	//open tracing
	//tracer, _, err2 := NewJaegerTracer("test")
	//if err2 != nil {
	//	grpclog.Errorf("new tracer err %v , continue", err2)
	//}
 
	opts = append(opts,grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
        grpc_ctxtags.StreamServerInterceptor(),
        //grpc_opentracing.StreamServerInterceptor(),
        //grpc_prometheus.StreamServerInterceptor,
        //grpc_zap.StreamServerInterceptor(zapLogger),
        grpc_auth.StreamServerInterceptor(AuthCheck),
        grpc_recovery.StreamServerInterceptor(),
    )))
	
	opts = append(opts,grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        grpc_ctxtags.UnaryServerInterceptor(),
        //grpc_opentracing.UnaryServerInterceptor(),
		//plugins.OpentracingServerInterceptor(tracer),
        //grpc_prometheus.UnaryServerInterceptor,
        //grpc_zap.UnaryServerInterceptor(zapLogger),
        grpc_auth.UnaryServerInterceptor(AuthCheck),
        grpc_recovery.UnaryServerInterceptor(),
		//Validate(),    //公共表单验证
    )))
 
    

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(grpcServer, newServer())
	
	//把server注册GRPC
	for _,inter := range ServerMap{
		fn:= reflect.ValueOf(inter["Register"])
		params := make([]reflect.Value, 2)
		params[0] = reflect.ValueOf(grpcServer)
		params[1] = reflect.ValueOf(inter["Server"])
		fn.Call(params)
	}
	
	
	//网关层
	
	//启动网关，把grpc注册HTTP
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	//mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	//	io.Copy(w, strings.NewReader(pb.Swagger))
	//})

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, demoAddr, dopts)
	
	//把server注册HTTP
	for _,inter := range ServerMap{
		fn:= reflect.ValueOf(inter["HandlerFromEndpoint"])
		params := make([]reflect.Value, 4)
		params[0] = reflect.ValueOf(ctx)
		params[1] = reflect.ValueOf(gwmux)
		params[2] = reflect.ValueOf(inter["Host"])
		params[3] = reflect.ValueOf(dopts)
		fn.Call(params)
	}
	
	
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	mux.Handle("/", gwmux)
	//serveSwagger(mux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", portHost))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("grpc on port: %d\n", portHost)
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//rb.Mq.RunConsumer(srv.ConsumersMap)
	return
}

//AuthCheck 鉴权验证
func AuthCheck(ctx context.Context) (context.Context,error){
	fmt.Println("login check function")
	return ctx,nil
}

//表单验证
func   Validate() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		out := reflect.ValueOf(req).MethodByName("Validate").Call(make([]reflect.Value, 0))
		if len(out)>0 {
			if out[0].Interface()!=nil {
				err = out[0].Interface().(error)
			}
		}
 		if   err!=nil {
 			return nil,err
		}
		resp, err = handler(ctx, req)
		return
	}
}

//NewJaegerTracer 初始化调用跟踪
func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "jaegertracing:6831",
		},
	}
	tracer, closer, err = cfg.New(
		serviceName,
		jaegerCfg.Logger(jaeger.StdLogger),
	)
	//defer closer.Close()

	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}