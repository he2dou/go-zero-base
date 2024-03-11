package customer

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"go-zero-base/app/user/internal/svc"
	"go-zero-base/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCustomerLogic {
	return &QueryCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCustomerLogic) QueryCustomer(req *types.QueryCustomerReq) (resp *types.QueryCustomerResp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.CustomerModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	resp = new(types.QueryCustomerResp)
	resp.Id = user.Id
	resp.NickName = user.UserName.String
	resp.Account = user.RealName.String
	resp.CreateTime = user.CreatedAt.Time.String()
	return
}
