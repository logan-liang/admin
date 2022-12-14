package svc

import (
	"admin/internal/config"

	"rpc/rpcclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc rpcclient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: rpcclient.NewRpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
