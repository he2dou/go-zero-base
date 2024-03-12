package customer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-base/app/money/internal/logic/customer"
	"go-zero-base/app/money/internal/svc"
	"go-zero-base/app/money/internal/types"
)

func DelCustomerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCustomerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewDelCustomerLogic(r.Context(), svcCtx)
		err := l.DelCustomer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
