package systemRoleMenu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleMenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleMenuDeleteLogic {
	return SystemRoleMenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleMenuDeleteLogic) SystemRoleMenuDelete(req types.SystemRoleMenuDelReq) (*types.SystemRoleMenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleMenuReply{}, nil
}
