package systemRoleApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleApiAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleApiAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleApiAddLogic {
	return SystemRoleApiAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleApiAddLogic) SystemRoleApiAdd(req types.RoleApisReq) (*types.SystemRoleApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleApiReply{}, nil
}
