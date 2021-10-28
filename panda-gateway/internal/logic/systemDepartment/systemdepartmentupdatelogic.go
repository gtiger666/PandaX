package systemDepartment

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentUpdateLogic {
	return SystemDepartmentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentUpdateLogic) SystemDepartmentUpdate(req types.SystemDepartmentPostReq) (*types.SystemDepartmentReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemDepartmentReply{}, nil
}
