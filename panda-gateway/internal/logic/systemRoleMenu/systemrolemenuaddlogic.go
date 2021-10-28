package systemRoleMenu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleMenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleMenuAddLogic {
	return SystemRoleMenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleMenuAddLogic) SystemRoleMenuAdd(req types.SystemRoleMenuAddReq) (*types.SystemRoleMenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleMenuReply{}, nil
}
