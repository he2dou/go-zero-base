package user

import (
	"context"

	"go-zero-base/app/money/internal/svc"
	"go-zero-base/app/money/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
