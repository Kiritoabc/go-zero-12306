package main

import (
	"flag"
	"fmt"
	"go-zero-12306/app/user-service/cmd/rpc/internal/config"
	"go-zero-12306/app/user-service/cmd/rpc/internal/server"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// configFile path
var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	// load config file
	conf.MustLoad(*configFile, &c)
	// new a service ctx
	ctx := svc.NewServiceContext(c)
	// create grpc service
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// register userService
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
