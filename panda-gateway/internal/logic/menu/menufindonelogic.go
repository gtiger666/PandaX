package menu

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuFindOneLogic {
	return MenuFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuFindOneLogic) MenuFindOne(req types.MenuDelReq) (*types.MenuReply, error) {
	// todo: add your logic here and delete this line

	return &types.MenuReply{}, nil
}
