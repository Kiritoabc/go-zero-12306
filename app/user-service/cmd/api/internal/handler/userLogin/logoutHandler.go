package userLogin

import (
	"fmt"
	"go-zero-12306/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-12306/app/user-service/cmd/api/internal/logic/userLogin"
	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		uid := r.Context().Value("jwtUserId")
		fmt.Sprintf("uid:%v", uid)
		l := userLogin.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)

		result.HttpResult(r, w, resp, err)
	}
}
