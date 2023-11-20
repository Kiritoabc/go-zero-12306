package tempSeat

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/logic/tempSeat"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"
)

func PurchaseTicketsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PurchaseTicketsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tempSeat.NewPurchaseTicketsLogic(r.Context(), svcCtx)
		resp, err := l.PurchaseTickets(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
