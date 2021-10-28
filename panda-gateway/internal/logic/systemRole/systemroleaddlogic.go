package systemRole

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleAddLogic {
	return SystemRoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleAddLogic) SystemRoleAdd(req types.SystemRolePostReq) (*types.SystemRoleReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleReply{}, nil
}
