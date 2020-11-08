package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	pb "micro/chapter4/examples/helloworld"
	"micro/mygrpc"
	"time"
)

func main() {
	cc, err := clientv3.New(clientv3.Config{
		Endpoints:        []string{"192.168.0.105:2379"},
		AutoSyncInterval: 0,
		DialTimeout:      Duration("1s"),
	})
	if err != nil {
		panic(err)
	}
	resolver.Register(mygrpc.NewResolver(cc))
	c := pb.NewGreeterClient(newGRPCClient())
	// Contact the server and print out its response.
	for i := 0; i < 40000; i++ {
		reply, _ := c.Token(context.Background(), &pb.TokenRequest{
			GrantType: "password",
			Username:  "test",
			Password:  "test",
			Scope:     "haha",
		})
		fmt.Println("reply------>", reply, i)
		time.Sleep(5 * time.Millisecond)
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
	cc, err := grpc.DialContext(ctx, "etcd:///micro", options...)

	if err != nil {
		panic(err)
	}
	return cc
}
