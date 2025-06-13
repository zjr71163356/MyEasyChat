package logic

import (
	"context"

	"myeasychat/user/api/internal/svc"
	"myeasychat/user/api/internal/types"
	"myeasychat/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetUserInfoResp{
		Id:    res.Id,
		Name:  res.Name,
		Phone: res.Phone,
	}
	return resp, nil
}
