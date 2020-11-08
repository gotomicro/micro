package main

import (
	"context"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	pb "micro/chapter4/examples/helloworld"
	"micro/mygrpc"
	"os"
)

var oauth2Server *server.Server

func main() {
	// 设置日志格式
	log.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","file":"${short_file}","line":"${line}"}`)
	// 全局日志级别
	log.SetLevel(log.DEBUG)
	log.Infof("server start, pid = %d", os.Getpid())

	var servOpts []grpc.ServerOption

	app := mygrpc.NewApp(
		mygrpc.WithServerName("micro"),
		mygrpc.WithGRPCServOption(servOpts),
	)
	app.Register(pb.RegisterGreeterServer, &Hello{})
	app.Start()
	log.Info("handle end")
}

type Hello struct{}

// SayHello implements helloworld.GreeterServer
func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof("receive req : %v", *in)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
