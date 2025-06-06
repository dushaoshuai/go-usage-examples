package main

import (
	"flag"
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/config"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/handler"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/swagger.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	handler.RegisterSwaggerHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
