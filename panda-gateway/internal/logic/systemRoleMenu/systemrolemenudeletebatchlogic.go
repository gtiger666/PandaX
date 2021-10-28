package systemRoleMenu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleMenuDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleMenuDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleMenuDeleteBatchLogic {
	return SystemRoleMenuDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleMenuDeleteBatchLogic) SystemRoleMenuDeleteBatch(req types.SystemRoleMenuDelBatchReq) (*types.SystemRoleMenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleMenuReply{}, nil
}
