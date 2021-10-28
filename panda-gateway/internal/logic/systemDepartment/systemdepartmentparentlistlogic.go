package systemDepartment

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentParentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentParentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentParentListLogic {
	return SystemDepartmentParentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentParentListLogic) SystemDepartmentParentList(req types.SystemDepartmentListReq) (*types.SystemDepartmentReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemDepartmentReply{}, nil
}
