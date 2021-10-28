package systemRoleApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleApiFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleApiFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleApiFindOneLogic {
	return SystemRoleApiFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleApiFindOneLogic) SystemRoleApiFindOne(req types.SystemRoleApiDelReq) (*types.SystemRoleApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleApiReply{}, nil
}
