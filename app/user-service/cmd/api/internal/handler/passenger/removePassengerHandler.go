package passenger

import (
	"go-zero-12306/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/user-service/cmd/api/internal/logic/passenger"
	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"
)

func RemovePassengerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemovePassengerReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := passenger.NewRemovePassengerLogic(r.Context(), svcCtx)
		resp, err := l.RemovePassenger(&req)
		result.HttpResult(r, w, resp, err)
	}
}
