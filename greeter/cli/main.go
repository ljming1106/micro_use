package main

import (
	"context"
	"fmt"
	"sync"

	client "github.com/micro/go-micro/v2/client"
	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro/v2"

	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
)

func main() {
	// create a new service
	service := micro.NewService(
		// wrap the client
		micro.WrapClient(logWrap),
		micro.WrapClient(hystrix.NewClientWrapper()),
		)

	// parse command line flags
	service.Init()

	// concurrent request
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	for i := 0;i<num;i++ {
		go func() {
			// Use the generated client stub
			cl := hello.NewSayService("go.micro.srv.greeter1", service.Client())
			// Make request
			rsp, err := cl.Hello(context.Background(), &hello.Request{
				Name: "John",
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(rsp.Msg)
			wg.Done()
		}()
	}
	wg.Wait()
}

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request service: %s method: %s\n", req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}