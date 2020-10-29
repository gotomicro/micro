package dns

import (
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/dns"
)

func RegisterResolver() {
	resolver.Register(&ResolverBuilder{})
}

type ResolverBuilder struct {
	closed bool
}

func (b *ResolverBuilder) ResolveNow(resolver.ResolveNowOption) {
}

func (b *ResolverBuilder) Close() {
	b.closed = true
}

func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	newResolver := dns.NewBuilder()
	return newResolver.Build(target, cc, opts)
}

func (b *ResolverBuilder) Scheme() string {
	return "dns"
}
