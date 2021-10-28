package systemRole

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleFindOneLogic {
	return SystemRoleFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleFindOneLogic) SystemRoleFindOne(req types.SystemRoleDelReq) (*types.SystemRoleReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleReply{}, nil
}
