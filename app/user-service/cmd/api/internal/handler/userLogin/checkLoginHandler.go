package userLogin

import (
	"go-zero-12306/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/user-service/cmd/api/internal/logic/userLogin"
	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"
)

func CheckLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := userLogin.NewCheckLoginLogic(r.Context(), svcCtx)
		resp, err := l.CheckLogin(&req)

		result.HttpResult(r, w, resp, err)
	}
}
