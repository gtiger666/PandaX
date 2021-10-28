package systemRoleMenu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleMenusByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleMenusByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleMenusByRoleIdLogic {
	return SystemRoleMenusByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleMenusByRoleIdLogic) SystemRoleMenusByRoleId(req types.SystemRoleMenuAddReq) (*types.SystemRoleMenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleMenuReply{}, nil
}
