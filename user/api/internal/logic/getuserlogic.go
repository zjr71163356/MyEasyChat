package logic

import (
	"context"

	"myeasychat/user/api/internal/svc"
	"myeasychat/user/api/internal/types"
	"myeasychat/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserResp{
		Id:    res.Id,
		Name:  res.Name,
		Phone: res.Phone,
	}
	return resp, nil
}
