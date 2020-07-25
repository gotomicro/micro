package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"micro/mygrpc"
	pb "micro/mygrpc/examples/helloworld"
	"time"
)

func main() {
	cc, err := clientv3.New(clientv3.Config{
		Endpoints:        []string{"127.0.0.1:2379"},
		AutoSyncInterval: 0,
		DialTimeout:      Duration("1s"),
	})
	if err != nil {
		panic(err)
	}
	resolver.Register(mygrpc.NewResolver(cc))
	c := pb.NewGreeterClient(newGRPCClient())
	// Contact the server and print out its response.
	name := "goto micro"

	rsp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Errorf("could not greet: %v", err)

	}
	log.Infof("Greeting: %s", rsp.Message)
	time.Sleep(time.Second * 2)
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
	cc, err := grpc.DialContext(ctx, "etcd:///micro", options...)

	if err != nil {
		panic(err)
	}
	return cc
}
