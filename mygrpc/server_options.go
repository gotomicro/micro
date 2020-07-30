package mygrpc

import (
	"google.golang.org/grpc"
)

//ServerOption option of server
type ServerOption struct {
	serverName string
	address    string
	registry   *Registry
	grpcOpts   []grpc.ServerOption
}

type ServerOptions func(o *ServerOption)

//WithRegistry set registry
func WithRegistry(r *Registry) ServerOptions {
	return func(o *ServerOption) {
		o.registry = r
	}
}

//WithGRPCServOption set grpc options
func WithGRPCServOption(opts []grpc.ServerOption) ServerOptions {
	return func(o *ServerOption) {
		o.grpcOpts = opts
	}
}

func WithServerName(sn string) ServerOptions {
	return func(o *ServerOption) {
		o.serverName = sn
	}
}

func WithAddress(address string) ServerOptions {
	return func(o *ServerOption) {
		o.address = address
	}
}
