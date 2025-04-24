package logic

import (
	"context"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JsonSimpleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonSimpleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonSimpleLogic {
	return &JsonSimpleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JsonSimpleLogic) JsonSimple(req *types.JsonReq) (resp *types.JsonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
