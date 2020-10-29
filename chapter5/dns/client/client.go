package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	pb "micro/chapter4/examples/helloworld"
	"os"
	"time"
)

func main() {
	// 设置日志格式
	log.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","file":"${short_file}","line":"${line}"}`)
	// 全局日志级别
	log.SetLevel(log.DEBUG)
	log.Infof("server start, pid = %d", os.Getpid())

	resolver.Register(resolver.Get("dns"))
	c := pb.NewGreeterClient(newGRPCClient())
	name := "goto micro"
	for {
		rsp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Errorf("could not greet: %v", err)
		} else {
			log.Infof("Greeting: %s", rsp.Message)
		}
		time.Sleep(time.Second * 2)
	}

}

func Duration(str string) time.Duration {
	dur, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}
	return dur
}

func newGRPCClient() *grpc.ClientConn {
	var ctx = context.Background()
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	cc, err := grpc.DialContext(ctx, "dns:///micro", options...)

	if err != nil {
		panic(err)
	}
	return cc
}
