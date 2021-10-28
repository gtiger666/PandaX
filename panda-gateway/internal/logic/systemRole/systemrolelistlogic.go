package systemRole

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleListLogic {
	return SystemRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleListLogic) SystemRoleList(req types.SystemRoleListReq) (*types.SystemRoleReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleReply{}, nil
}
