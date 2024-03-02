package regionStation

import (
	"go-zero-12306/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/ticket-service/cmd/api/internal/logic/regionStation"
	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"
)

func ListRegionStationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegionStationQueryReqDTO
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := regionStation.NewListRegionStationLogic(r.Context(), svcCtx)
		resp, err := l.ListRegionStation(&req)
		result.HttpResult(r, w, resp, err)
	}
}
