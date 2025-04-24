package handler

import (
	"net/http"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/logic"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func jsonComplexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComplexJsonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJsonComplexLogic(r.Context(), svcCtx)
		resp, err := l.JsonComplex(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
