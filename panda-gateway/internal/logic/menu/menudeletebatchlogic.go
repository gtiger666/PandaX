package menu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuDeleteBatchLogic {
	return MenuDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteBatchLogic) MenuDeleteBatch(req types.MenuDelBatchReq) (*types.MenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.MenuReply{}, nil
}
