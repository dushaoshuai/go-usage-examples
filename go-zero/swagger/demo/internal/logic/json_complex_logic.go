package logic

import (
	"context"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JsonComplexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonComplexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonComplexLogic {
	return &JsonComplexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JsonComplexLogic) JsonComplex(req *types.ComplexJsonReq) (resp *types.ComplexJsonResp, err error) {
	return &types.ComplexJsonResp{
		Integer:      1,
		Number:       1,
		Boolean:      true,
		String:       "kitty",
		ArrayInteger: []int{1, 2, 3},
	}, nil
}
