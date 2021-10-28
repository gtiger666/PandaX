package systemApi

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemApiListLogic {
	return SystemApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemApiListLogic) SystemApiList(req types.SystemApiListReq) (*types.SystemApiReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemApiReply{}, nil
}
