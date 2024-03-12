package customer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-base/app/money/internal/logic/customer"
	"go-zero-base/app/money/internal/svc"
	"go-zero-base/app/money/internal/types"
)

func AddCustomerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCustomerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewAddCustomerLogic(r.Context(), svcCtx)
		resp, err := l.AddCustomer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
