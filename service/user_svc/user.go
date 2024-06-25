package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"zero-mall/service/user_svc/internal/config"
	"zero-mall/service/user_svc/internal/server"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting user_svc server at %s...\n", c.ListenOn)
	s.Start()
}
