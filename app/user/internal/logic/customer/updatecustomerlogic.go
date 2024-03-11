package customer

import (
	"context"
	"database/sql"
	"go-zero-base/app/user/internal/model/customer"

	"go-zero-base/app/user/internal/svc"
	"go-zero-base/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(req *types.UpdateCustomerReq) error {
	// todo: add your logic here and delete this line
	c := customer.Customer{
		Id: req.Id,
	}
	var flag bool
	flag = len(req.UserName) > 0
	c.UserName = sql.NullString{
		String: req.UserName,
		Valid:  flag,
	}
	flag = len(req.Industry) > 0
	c.Industry = sql.NullString{
		String: req.Industry,
		Valid:  flag,
	}
	l.svcCtx.CustomerModel.Update(l.ctx, nil, &c)
	return nil
}
