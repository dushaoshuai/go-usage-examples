package handler

import (
	"net/http"

	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/logic"
	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
