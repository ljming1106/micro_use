// Package main
package main

import (
	"context"
	"github.com/micro/examples/greeter/common/tracer"

	"time"

	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro/v2"
	server "github.com/micro/go-micro/v2/server"
	//"github.com/micro/go-micro/v2/util/log"
	log "github.com/micro/examples/greeter/common/logger"

	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"google.golang.org/grpc"
	// 引入插件
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

// MacOS run：
// CC=x86_64-linux-musl-gcc  GOARCH=amd64 CGO_ENABLED=0 go run  main.go

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	//time.Sleep(800 * time.Millisecond)
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

const QPS = 5

// 所有port都是服务默认端口，ip请根据实际情况配置
const (
	// 项目服务名和注册到Jaeger的服务名一致，保障更容易定位到具体服务
	ServerName = "go.micro.srv.greeter1"
	// 服务地址
	JaegerAddr = "127.0.0.1:6831"
)

func main() {
	log.Info("server start...")
	go func() {
		for {
			grpc.DialContext(context.TODO(), "127.0.0.1:9091")
			time.Sleep(time.Second)
		}
	}()

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServerName, JaegerAddr)
	if err != nil {
		log.Panicf("NewJaegerTracer fail!%v", err)
	}
	defer closer.Close()

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter1"),
		// wrap the handler
		micro.WrapHandler(logWrapper),
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),
		// 配置链路追踪为jaeger
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	// optionally setup command line usage
	log.Info("service init...")
	service.Init()

	log.Info("register handlers...")
	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	log.Info("service run...")
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Infof("[wrapper] server request: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		return err
	}
}
