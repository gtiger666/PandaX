package admin

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminDeleteLogic {
	return AdminDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminDeleteLogic) AdminDelete(req types.AdminDelReq) (*types.AdminReply, error) {
	// todo: add your logic here and delete this line

	return &types.AdminReply{}, nil
}
