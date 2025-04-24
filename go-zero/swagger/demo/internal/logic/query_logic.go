package logic

import (
	"context"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.QueryReq) (resp *types.QueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
