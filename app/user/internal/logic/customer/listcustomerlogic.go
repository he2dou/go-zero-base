package customer

import (
	"context"
	"go-zero-base/app/user/internal/svc"
	"go-zero-base/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCustomerLogic {
	return &ListCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCustomerLogic) ListCustomer(req *types.ListCustomerReq) (resp *types.ListCustomerResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ListCustomerResp)

	total, err := l.svcCtx.CustomerModel.FindPage(l.ctx, nil, req, &resp.List)
	if err != nil {
		return
	}
	resp.PageInfo.Total = total
	return
}
