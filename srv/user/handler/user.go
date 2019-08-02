package handler

import (
	"fmt"

	"github.com/mouuii/go-easy-blog/common/store"
	"github.com/mouuii/go-easy-blog/srv/user/proto/user"

	"golang.org/x/net/context"
)

type UserService struct {
	db *store.DB
}

func NewUserService(db *store.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Login(ctx context.Context, req *user.LoginReq, resp *user.LoginReq) error {
	fmt.Printf("%#v\n", *req)
	*resp = *req
	return nil
}
