package main

import (
	"log"
	"time"

	"github.com/mouuii/go-easy-blog/api/user/handler"
	"github.com/mouuii/go-easy-blog/api/user/router"
	"github.com/mouuii/go-easy-blog/common"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
)

var version = "0.0.1"

func main() {

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := web.NewService(
		web.Name(common.ApiUserName),
		web.Version(version),
		web.RegisterInterval(time.Second*2),
		web.Registry(reg),
	)

	srv.Init()

	r := router.InitRouter()
	srv.Handle("/", r)

	handler.RegisterClient(common.Client())

	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}
}
