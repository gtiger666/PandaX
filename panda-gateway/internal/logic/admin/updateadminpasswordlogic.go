package admin

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateadminpasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateadminpasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateadminpasswordLogic {
	return UpdateadminpasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateadminpasswordLogic) Updateadminpassword(req types.AdminUpdatePwdReq) (*types.AdminReply, error) {
	// todo: add your logic here and delete this line

	return &types.AdminReply{}, nil
}
