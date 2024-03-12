package customer

import (
	"context"

	"go-zero-base/app/money/internal/svc"
	"go-zero-base/app/money/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCustomerLogic {
	return &DelCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCustomerLogic) DelCustomer(req *types.DeleteCustomerReq) error {
	// todo: add your logic here and delete this line

	return nil
}
