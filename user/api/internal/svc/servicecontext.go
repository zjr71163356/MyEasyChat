package svc

import (
	"myeasychat/user/api/internal/config"
	"myeasychat/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
