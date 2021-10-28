package admin

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminDeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminDeleteBatchLogic {
	return AdminDeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminDeleteBatchLogic) AdminDeleteBatch(req types.AdminDelBatchReq) (*types.AdminReply, error) {
	// todo: add your logic here and delete this line

	return &types.AdminReply{}, nil
}
