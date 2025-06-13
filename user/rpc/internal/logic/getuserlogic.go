package logic

import (
	"context"

	"myeasychat/user/rpc/internal/svc"
	"myeasychat/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResp{
		Id:    resp.Id,
		Name:  resp.Name,
		Phone: resp.Phone,
	}, nil
}
