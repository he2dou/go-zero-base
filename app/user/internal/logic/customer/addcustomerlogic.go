package customer

import (
	"context"

	"go-zero-base/app/user/internal/svc"
	"go-zero-base/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCustomerLogic {
	return &AddCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCustomerLogic) AddCustomer(req *types.AddCustomerReq) (resp *types.AddCustomerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
