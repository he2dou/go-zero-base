package user

import (
	"context"

	"go-zero-base/app/money/internal/svc"
	"go-zero-base/app/money/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
	return &QueryUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserLogic) QueryUser(req *types.QueryUserReq) (resp *types.QueryUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
