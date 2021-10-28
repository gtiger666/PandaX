package systemDepartment

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentFindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentFindOneLogic {
	return SystemDepartmentFindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentFindOneLogic) SystemDepartmentFindOne(req types.SystemDepartmentDelReq) (*types.SystemDepartmentReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemDepartmentReply{}, nil
}
