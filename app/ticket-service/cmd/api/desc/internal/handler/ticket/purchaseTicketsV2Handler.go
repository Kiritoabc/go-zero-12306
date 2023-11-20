package ticket

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/logic/ticket"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"
)

func PurchaseTicketsV2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PurchaseTicketReqDTO
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ticket.NewPurchaseTicketsV2Logic(r.Context(), svcCtx)
		resp, err := l.PurchaseTicketsV2(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
