package logic

import (
	"context"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPathLogic {
	return &QueryPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPathLogic) QueryPath(req *types.PathQueryReq) (resp *types.PathQueryResp, err error) {
	return &types.PathQueryResp{
		Id:   11,
		Name: "kitty",
	}, nil
}
