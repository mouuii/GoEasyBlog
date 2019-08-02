package handler

import (
	"github.com/mouuii/go-easy-blog/common/store"
	"github.com/mouuii/go-easy-blog/srv/user/proto/user"

	"github.com/micro/go-micro/server"
)

func InitService(srv server.Server, db *store.DB) {
	user.RegisterUserServiceHandler(srv, NewUserService(db))
}
