package main

import (
	"flag"
	"fmt"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/config"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/handler"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	_ "github.com/dtm-labs/driver-gozero"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the etc file")

func init() {
	//close statis log
	//logx.DisableStat()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 注册所有的 rpc client 服务, 获取 ServiceContext
	ctx := svc.NewServiceContext(c)

	// 注册当前服务， http server
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
