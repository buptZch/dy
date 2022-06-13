package main

import (
	"dy/cmd/user_action/dal"
	useraction "dy/cmd/user_action/kitex_gen/useraction/useractionservice"
	"dy/pkg/bound"
	"dy/pkg/constants"
	"dy/pkg/middleware"
	tracer2 "dy/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.UserActionServiceName)
	dal.Init()
}

func main() {
	//首先生成Etcd的注册对象
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	//监听的本地ip
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	//初始化
	Init()
	svr := useraction.NewServer(new(UserActionServiceImpl),
		//指定服务名
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserActionServiceName}), // server name
		//传递mideleware
		server.WithMiddleware(middleware.CommonMiddleware), // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		//传递服务地址
		server.WithServiceAddr(addr), // address
		//传递限流的参数
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//选择连接方式 https://www.cloudwego.io/zh/docs/kitex/tutorials/basic-feature/connection_type/
		//连接多路复用
		server.WithMuxTransport(), // Multiplex
		//先在init中初始化jaeger
		server.WithSuite(trace.NewDefaultServerSuite()), // tracer
		//cpu求情级别的限流器
		server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r),                              // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
