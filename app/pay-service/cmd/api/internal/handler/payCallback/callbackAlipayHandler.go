package payCallback

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/pay-service/cmd/api/internal/logic/payCallback"
	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"
)

func CallbackAlipayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackAlipayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := payCallback.NewCallbackAlipayLogic(r.Context(), svcCtx)
		resp, err := l.CallbackAlipay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
