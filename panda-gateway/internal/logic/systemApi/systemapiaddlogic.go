package systemApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemApiAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemApiAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemApiAddLogic {
	return SystemApiAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemApiAddLogic) SystemApiAdd(req types.SystemApiPostReq) (*types.SystemApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemApiReply{}, nil
}
