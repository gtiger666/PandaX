package systemRoleMenu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleMenuFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleMenuFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleMenuFindOneLogic {
	return SystemRoleMenuFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleMenuFindOneLogic) SystemRoleMenuFindOne(req types.SystemRoleMenuDelReq) (*types.SystemRoleMenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleMenuReply{}, nil
}
