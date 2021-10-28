package systemApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemApiDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemApiDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemApiDeleteLogic {
	return SystemApiDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemApiDeleteLogic) SystemApiDelete(req types.SystemApiDelReq) (*types.SystemApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemApiReply{}, nil
}
