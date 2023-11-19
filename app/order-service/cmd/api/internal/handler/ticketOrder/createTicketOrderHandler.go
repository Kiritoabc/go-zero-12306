package ticketOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/order-service/cmd/api/internal/logic/ticketOrder"
	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"
)

func CreateTicketOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTicketOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ticketOrder.NewCreateTicketOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateTicketOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
