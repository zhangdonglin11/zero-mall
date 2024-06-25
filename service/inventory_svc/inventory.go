package main

import (
	"flag"
	"fmt"
	"zero-mall/service/inventory_svc/internal/config"
	"zero-mall/service/inventory_svc/internal/server"
	"zero-mall/service/inventory_svc/internal/svc"
	"zero-mall/service/inventory_svc/types/inventory"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/inventory.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		inventory.RegisterInventoryServer(grpcServer, server.NewInventoryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting user_svc server at %s...\n", c.ListenOn)
	s.Start()
}
