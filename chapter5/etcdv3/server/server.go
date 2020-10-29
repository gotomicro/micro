package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	pb "micro/chapter3/examples/helloworld"
	"micro/mygrpc"
	"os"
	"time"
)

func main() {
	log.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","file":"${short_file}","line":"${line}"}`) // 设置日志格式
	log.SetLevel(log.DEBUG)                                                                                // 全局日志级别
	log.Infof("server start, pid = %d", os.Getpid())

	cc, err := clientv3.New(clientv3.Config{
		Endpoints:        []string{"127.0.0.1:2379"}, // etcd节点ip
		AutoSyncInterval: Duration("60s"),            // 自动同步etcd的member节点
		DialTimeout:      Duration("1s"),             // 拨号超时时间
	})
	if err != nil {
		panic(err)
	}

	var servOpts []grpc.ServerOption

	app := mygrpc.NewApp(
		mygrpc.WithAddress("127.0.0.1:4000"),       // 设置服务Address
		mygrpc.WithRegistry(mygrpc.NewRegisty(cc)), // 设置服务注册中心
		mygrpc.WithServerName("micro"),             // 设置服务名称
		mygrpc.WithGRPCServOption(servOpts),        // 设置服务属性
	)
	app.Register(pb.RegisterGreeterServer, &Hello{})
	app.Start()
	log.Info("handle end")
}

func Duration(str string) time.Duration {
	dur, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}
	return dur
}

type Hello struct{}

// SayHello implements helloworld.GreeterServer
func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof("receive req : %v", *in)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
