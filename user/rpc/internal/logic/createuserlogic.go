package logic

import (
	"context"
	"database/sql"

	"myeasychat/user/models"
	"myeasychat/user/rpc/internal/svc"
	"myeasychat/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	// todo: add your logic here and delete this line
	Resp, err := l.svcCtx.UsersModel.Insert(l.ctx, &models.Users{
		Name:     in.Name,
		Phone:    in.Phone,
		Password: sql.NullString{String: in.Password, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	newId, err := Resp.LastInsertId()
	return &user.CreateUserResp{
		Id: string(newId),
	}, nil
}
