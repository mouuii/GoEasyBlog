package common

import (
	"time"

	"github.com/micro/go-micro/client"
)

var defaultTimeOut = time.Second * 360


func Client() client.Client {
	return client.NewClient(client.RequestTimeout(defaultTimeOut))
}


func ClientWithTimeOut(second time.Duration) client.Client {
	return client.NewClient(client.RequestTimeout(second))
}
