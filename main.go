package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"grpc-balancer/pb"
	"grpc-balancer/service"
	"net"
	"strconv"
	"time"
)

func main() {
	go runServer(7788, "服务1")
	go runServer(7789, "服务2")
	go runServer(7799, "服务3")
	call := rpcCall(NewBuilder([]string{"localhost:7788",
		"localhost:7789", "localhost:7799"}))
	for {
		svc, err := call.GetSvc(context.Background(), &pb.GetSvcRequest{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(svc.Result)
		}
		time.Sleep(2 * time.Second)
	}
}

func runServer(port int, tag string) {
	options := make([]grpc.ServerOption, 0)
	server := grpc.NewServer(options...)
	pb.RegisterSvcServer(server, service.NewSvcService(tag))
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func rpcCall(builder resolver.Builder) pb.SvcClient {
	options := make([]grpc.DialOption, 0)
	options = append(options, grpc.WithTransportCredentials(insecure.
		NewCredentials()))
	// 轮询策略，修改规则 https://github.com/grpc/grpc/blob/master/doc/service_config.md
	poli := `{"loadBalancingPolicy":"round_robin"}`
	options = append(options, grpc.WithDefaultServiceConfig(poli))
	// 自定义服务地址解析器，提供所有可用的服务列表
	options = append(options, grpc.WithResolvers(builder))
	con, err := grpc.Dial(builder.Scheme()+"://0.0.0.0", options...)
	if err != nil {
		panic(err)
	}
	return pb.NewSvcClient(con)
}
