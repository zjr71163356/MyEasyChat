package logic

import (
	"context"

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
	_, err := l.svcCtx.UsersModel.Insert(l.ctx, &models.Users{
		Id:    in.Id,
		Name:  in.Name,
		Phone: in.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &user.CreateUserResp{}, nil
}
