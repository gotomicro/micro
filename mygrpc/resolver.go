package mygrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

const (
	resolverTimeOut = 10 * time.Second
)

type Resolver struct {
	cli *clientv3.Client
}

//NewResolver create a resolver for grpc
func NewResolver(cli *clientv3.Client) *Resolver {
	return &Resolver{
		cli: cli,
	}
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	go r.watch(cc, target.Endpoint)
	return r, nil
}

func (r *Resolver) Scheme() string {
	return "etcd"
}

// ResolveNow ...
func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

// close closes the resolver.
func (r *Resolver) Close() {

}

func (r *Resolver) watch(cc resolver.ClientConn, serviceName string) {
	target := fmt.Sprintf("/%s/", serviceName)
	for {
		resolverObj := NewAddressList()
		resp, err := r.cli.Get(context.Background(), target, clientv3.WithPrefix())
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}
		// 初始化
		for _, value := range resp.Kvs {
			var resolverInfo resolver.Address
			if err := json.Unmarshal(value.Value, &resolverInfo); err != nil {
				continue
			}
			resolverObj.Put(resolverInfo)
		}

		cc.UpdateState(resolver.State{
			Addresses: resolverObj.GetAddressList(),
		})

		// watch
		ctx, cancel := context.WithCancel(context.Background())
		rch := r.cli.Watch(ctx, target, clientv3.WithPrefix(), clientv3.WithCreatedNotify())
		for n := range rch {
			for _, ev := range n.Events {
				switch ev.Type {
				// 添加或者更新
				case mvccpb.PUT:
					var resolverInfo resolver.Address
					if err := json.Unmarshal(ev.Kv.Value, &resolverInfo); err == nil {
						resolverObj.Put(resolverInfo)
					}

				// 硬删除
				case mvccpb.DELETE:
					var resolverInfo resolver.Address
					if err := json.Unmarshal(ev.Kv.Value, &resolverInfo); err == nil {
						resolverObj.Delete(resolverInfo)
					}
				}
			}
			cc.UpdateState(resolver.State{
				Addresses: resolverObj.GetAddressList(),
			})
		}
		cancel()
	}
}

type AddressList struct {
	serverName string
	store      map[string]resolver.Address
	m          sync.RWMutex
}

func NewAddressList() *AddressList {
	return &AddressList{
		store: make(map[string]resolver.Address),
	}
}

func (a *AddressList) Put(address resolver.Address) {
	a.m.Lock()
	defer a.m.Unlock()
	a.store[address.Addr] = address
}

func (a *AddressList) Delete(address resolver.Address) {
	a.m.Lock()
	defer a.m.Unlock()
	delete(a.store, address.Addr)
}

func (a *AddressList) GetAddressList() []resolver.Address {
	addrs := make([]resolver.Address, 0)
	a.m.RLock()
	defer a.m.RUnlock()
	for _, address := range a.store {
		addrs = append(addrs, address)
	}
	return addrs
}
