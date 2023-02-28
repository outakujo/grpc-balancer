package main

import (
	"google.golang.org/grpc/resolver"
)

type Builder struct {
	ads []string
}

func NewBuilder(ads []string) *Builder {
	return &Builder{ads: ads}
}

// 返回解析后可用的服务地址列表
func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	adds := make([]resolver.Address, 0)
	for _, ad := range b.ads {
		adds = append(adds, resolver.Address{Addr: ad})
	}
	state := resolver.State{Addresses: adds}
	err := cc.UpdateState(state)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// 协议自定义即可，在build方法中target使用
func (b *Builder) Scheme() string {
	return "lb"
}

func (b *Builder) ResolveNow(options resolver.ResolveNowOptions) {

}

func (b *Builder) Close() {

}
