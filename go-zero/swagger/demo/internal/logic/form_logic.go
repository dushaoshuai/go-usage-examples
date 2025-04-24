package logic

import (
	"context"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FormLogic {
	return &FormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FormLogic) Form(req *types.FormReq) (resp *types.FormResp, err error) {
	// todo: add your logic here and delete this line

	return
}
