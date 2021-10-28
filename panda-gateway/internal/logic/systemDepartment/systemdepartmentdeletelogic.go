package systemDepartment

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentDeleteLogic {
	return SystemDepartmentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentDeleteLogic) SystemDepartmentDelete(req types.SystemDepartmentDelReq) (*types.SystemDepartmentReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemDepartmentReply{}, nil
}
