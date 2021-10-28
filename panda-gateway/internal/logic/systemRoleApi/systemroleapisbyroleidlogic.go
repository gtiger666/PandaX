package systemRoleApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleApisByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleApisByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleApisByRoleIdLogic {
	return SystemRoleApisByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleApisByRoleIdLogic) SystemRoleApisByRoleId(req types.SystemRoleApiDelReq) (*types.SystemRoleApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleApiReply{}, nil
}
