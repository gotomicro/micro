package mygrpc

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"time"
)

type Registry struct {
	cancel context.CancelFunc
	cli    *clientv3.Client
	lsCli  clientv3.Lease
}

//DefaultRegInfTTL default ttl of server info in registry
const DefaultRegInfTTL = time.Second * 50
const resolverTimeOut = time.Second * 2

type RegistryOption struct {
	TTL time.Duration
}
type RegistryOptions func(o *RegistryOption)

//NewRegisty create a reistry for registering server addr
func NewRegisty(cli *clientv3.Client) *Registry {
	return &Registry{
		cli:   cli,
		lsCli: clientv3.NewLease(cli),
	}
}

func (er *Registry) Register(ctx context.Context, serverName string, addr string, opts ...RegistryOptions) (err error) {
	var upBytes []byte
	info := resolver.Address{
		Addr:       addr,
		ServerName: serverName,
		Attributes: nil,
	}

	if upBytes, err = json.Marshal(info); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.TODO(), resolverTimeOut)
	er.cancel = cancel
	rgOpt := RegistryOption{TTL: DefaultRegInfTTL}
	for _, opt := range opts {
		opt(&rgOpt)
	}
	key := "/" + serverName + "/" + addr

	lsRsp, err := er.lsCli.Grant(ctx, int64(rgOpt.TTL/time.Second))
	if err != nil {
		return err
	}
	etcdOpts := []clientv3.OpOption{clientv3.WithLease(lsRsp.ID)}
	_, err = er.cli.KV.Put(ctx, key, string(upBytes), etcdOpts...)
	if err != nil {
		return err
	}
	lsRspChan, err := er.lsCli.KeepAlive(context.TODO(), lsRsp.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			_, ok := <-lsRspChan
			if !ok {
				grpclog.Fatalf("%v keepalive channel is closing", key)
				break
			}
		}
	}()
	return nil
}

func (er *Registry) Unregister(ctx context.Context, serverName string, addr string) (err error) {
	key := "/" + serverName + "/" + addr
	_, err = er.cli.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (er *Registry) Close() {
	er.cancel()
	er.cli.Close()
}
