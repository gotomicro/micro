package main

import (
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/labstack/gommon/log"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	pb "micro/chapter4/examples/helloworld"
	"micro/mygrpc"
	//"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"os"
	"time"
)

var oauth2Server *server.Server

func main() {
	// 设置日志格式
	log.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","file":"${short_file}","line":"${line}"}`)
	// 全局日志级别
	log.SetLevel(log.DEBUG)
	log.Infof("server start, pid = %d", os.Getpid())

	cc, err := clientv3.New(clientv3.Config{
		Endpoints:        []string{"127.0.0.1:2379"},
		AutoSyncInterval: 0,
		DialTimeout:      Duration("1s"),
	})
	if err != nil {
		panic(err)
	}
	initOauth2()

	builder, err := NewGlobalRateLimiterBuilder(func(i interface{}, handler grpc.UnaryHandler) (interface{}, error) {
		errMsg := " the request is rejected because of over QPS limitation"
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	})

	servOpts := []grpc.ServerOption{ grpc.UnaryInterceptor(builder.GlobalRateLimit)}

	app := mygrpc.NewApp(
		mygrpc.WithAddress("127.0.0.1:4000"),
		mygrpc.WithRegistry(mygrpc.NewRegisty(cc)),
		mygrpc.WithServerName("micro"),
		mygrpc.WithGRPCServOption(servOpts),
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

func (s *Hello) Token(ctx context.Context, in *pb.TokenRequest) (*pb.TokenReply, error) {
	userId, err := oauth2Server.PasswordAuthorizationHandler(in.Username, in.Password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error user")
	}
	fmt.Println(userId)
	return &pb.TokenReply{
		AccessToken: "success token",
		TokenType:   "test",
		ExpiresIn:   1800,
	}, nil
}

func initOauth2() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	//clientStore := store.NewClientStore()
	//clientStore.Set("000000", &models.Client{
	//	ID:     "000000",
	//	Secret: "999999",
	//	Domain: "http://localhost",
	//})
	//manager.MapClientStorage(clientStore)

	oauth2Server = server.NewDefaultServer(manager)
	oauth2Server.SetAllowGetAccessRequest(true)
	oauth2Server.SetClientInfoHandler(server.ClientFormHandler)

	oauth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		fmt.Println("Internal Error:", err.Error())
		return
	})

	oauth2Server.SetResponseErrorHandler(func(re *errors.Response) {
		fmt.Println("Response Error:", re.Error.Error())
	})

	oauth2Server.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
}
