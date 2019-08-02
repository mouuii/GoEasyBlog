package handler

import (
	"github.com/mouuii/go-easy-blog/srv/user/proto/user"
	"github.com/mouuii/go-easy-blog/common"
	
	client "github.com/micro/go-micro/client"
)

var (
	userClient user.UserServiceClient
)

func RegisterClient(cli client.Client) {
	userClient = user.NewUserServiceClient(common.SrvUserName, cli)
}