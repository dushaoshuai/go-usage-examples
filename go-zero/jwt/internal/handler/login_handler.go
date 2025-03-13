package handler

import (
	"net/http"

	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/logic"
	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
