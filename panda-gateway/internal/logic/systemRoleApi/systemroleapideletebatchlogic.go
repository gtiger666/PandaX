package systemRoleApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleApiDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleApiDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleApiDeleteBatchLogic {
	return SystemRoleApiDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleApiDeleteBatchLogic) SystemRoleApiDeleteBatch(req types.SystemRoleApiDelBatchReq) (*types.SystemRoleApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleApiReply{}, nil
}
