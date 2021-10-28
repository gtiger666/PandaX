package admin

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminFindOneLogic {
	return AdminFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminFindOneLogic) AdminFindOne(req types.AdminDelReq) (*types.AdminReply, error) {
	// todo: add your logic here and delete this line

	return &types.AdminReply{}, nil
}
