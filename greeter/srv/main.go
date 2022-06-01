// Package main
package main

import (
	"context"
	"fmt"
	"time"

	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	hello "github.com/micro/examples/greeter/srv/proto/hello"
	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"google.golang.org/grpc"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	time.Sleep(500*time.Millisecond)
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

const QPS = 5

func main() {
	go func() {
		for {
			grpc.DialContext(context.TODO(), "127.0.0.1:9091")
			time.Sleep(time.Second)
		}
	}()

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter1"),
		// wrap the handler
		micro.WrapHandler(logWrapper),
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[wrapper] server request: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		return err
	}
}