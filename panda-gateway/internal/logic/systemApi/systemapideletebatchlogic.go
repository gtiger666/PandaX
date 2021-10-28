package systemApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemApiDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemApiDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemApiDeleteBatchLogic {
	return SystemApiDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemApiDeleteBatchLogic) SystemApiDeleteBatch(req types.SystemApiDelBatchReq) (*types.SystemApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemApiReply{}, nil
}
