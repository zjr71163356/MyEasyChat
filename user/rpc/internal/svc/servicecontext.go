package svc

import (
	"myeasychat/user/models"
	"myeasychat/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUsersModel(sqlconn, c.Cache),
	}
}
