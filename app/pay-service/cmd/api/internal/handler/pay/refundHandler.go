package pay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/pay-service/cmd/api/internal/logic/pay"
	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"
)

func RefundHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefundReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pay.NewRefundLogic(r.Context(), svcCtx)
		resp, err := l.Refund(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
