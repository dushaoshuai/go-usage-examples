// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package handler

import (
	"net/http"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/query",
				Handler: queryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/query/:id",
				Handler: queryPathHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/form",
				Handler: formHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/json/complex",
				Handler: jsonComplexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/json/simple",
				Handler: jsonSimpleHandler(serverCtx),
			},
		},
	)
}
