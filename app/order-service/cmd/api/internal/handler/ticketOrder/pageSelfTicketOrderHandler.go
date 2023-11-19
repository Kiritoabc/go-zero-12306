package ticketOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/order-service/cmd/api/internal/logic/ticketOrder"
	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"
)

func PageSelfTicketOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TicketOrderSelfPageQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ticketOrder.NewPageSelfTicketOrderLogic(r.Context(), svcCtx)
		resp, err := l.PageSelfTicketOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
