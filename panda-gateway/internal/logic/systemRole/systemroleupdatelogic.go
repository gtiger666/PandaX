package systemRole

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleUpdateLogic {
	return SystemRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleUpdateLogic) SystemRoleUpdate(req types.SystemRolePostReq) (*types.SystemRoleReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleReply{}, nil
}
