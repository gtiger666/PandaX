package systemRoleApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRoleApiDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemRoleApiDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemRoleApiDeleteLogic {
	return SystemRoleApiDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemRoleApiDeleteLogic) SystemRoleApiDelete(req types.SystemRoleApiDelReq) (*types.SystemRoleApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemRoleApiReply{}, nil
}
