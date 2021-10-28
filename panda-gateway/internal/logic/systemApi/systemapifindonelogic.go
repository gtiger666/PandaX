package systemApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemApiFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemApiFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemApiFindOneLogic {
	return SystemApiFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemApiFindOneLogic) SystemApiFindOne(req types.SystemApiDelReq) (*types.SystemApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemApiReply{}, nil
}
