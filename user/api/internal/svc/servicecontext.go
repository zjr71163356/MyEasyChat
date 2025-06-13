package svc

import (
	"myeasychat/user/api/internal/config"
	"myeasychat/user/api/internal/middleware"
	"myeasychat/user/rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	// fmt.Println(c.UserRpc)

	return &ServiceContext{
		Config:            c,
		User:              userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
