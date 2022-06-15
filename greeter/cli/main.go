package main

import (
	"context"
	"github.com/micro/examples/greeter/common/tracer"
	"github.com/micro/go-micro/v2"
	//"github.com/micro/go-micro/v2/util/log"
	log "github.com/micro/examples/greeter/common/logger"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"

	"sync"

	hello "github.com/micro/examples/greeter/srv/proto/hello"
	client "github.com/micro/go-micro/v2/client"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

// MacOS run：
// CC=x86_64-linux-musl-gcc  GOARCH=amd64 CGO_ENABLED=0 go run  main.go

const (
	// 项目服务名和注册到Jaeger的服务名一致，保障更容易定位到具体服务
	ServerName = "go.micro.cli.greeter1"
	// 服务地址
	JaegerAddr = "127.0.0.1:6831"
)

func main() {
	log.Info("cli start...")
	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServerName, JaegerAddr)
	if err != nil {
		log.Panicf("NewJaegerTracer fail!%v", err)
	}
	defer closer.Close()

	// create a new service
	service := micro.NewService(
		// wrap the client
		micro.WrapClient(logWrap),
		micro.WrapClient(hystrix.NewClientWrapper(),
			// 配置链路追踪为jaeger
			opentracing.NewClientWrapper(jaegerTracer),
		),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379（docker对外映射端口12379）
		micro.Registry(etcd.NewRegistry(
			// 地址是我本地etcd服务器地址，不要照抄
			registry.Addrs("127.0.0.1:12379"),
		)),
	)

	log.Info("service init...")
	// parse command line flags
	service.Init()

	// concurrent request
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			// Use the generated client stub
			cl := hello.NewSayService("go.micro.srv.greeter1", service.Client())
			log.Infof("cli Hello:%v", i)
			// Make request
			rsp, err := cl.Hello(context.Background(), &hello.Request{
				Name: "John",
			})
			if err != nil {
				log.Info(err)
				return
			}
			log.Info(rsp.Msg)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Infof("cli finish!!!")
}

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Infof("[wrapper] client request service: %s method: %s\n", req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}
