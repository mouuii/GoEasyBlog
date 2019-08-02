package main

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/mouuii/go-easy-blog/common"
	"github.com/mouuii/go-easy-blog/srv/user/handler"
)

var version = "0.1.0"

func main() {
	srv := micro.NewService(
		micro.Name(common.SrvUserName),
		micro.Version(version),
		micro.RegisterInterval(time.Second*2),
		// micro.WrapHandler(middlewares.RPCCallLog),
	)

	srv.Init()

	handler.InitService(srv.Server(), nil)

	srv.Run()
}
