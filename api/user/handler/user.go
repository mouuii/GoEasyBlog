package handler

import (
	"fmt"

	"github.com/mouuii/go-easy-blog/srv/user/proto/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func Login(c *gin.Context) {
	req := user.LoginReq{}

	fmt.Println(c.Bind(&req))

	resp, _ := userClient.Login(context.Background(), &req)

	fmt.Println(*resp)
	c.JSON(200, resp)
}
