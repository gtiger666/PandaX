package systemRole

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleDeleteLogic {
	return SystemRoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleDeleteLogic) SystemRoleDelete(req types.SystemRoleDelReq) (*types.SystemRoleReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleReply{}, nil
}
