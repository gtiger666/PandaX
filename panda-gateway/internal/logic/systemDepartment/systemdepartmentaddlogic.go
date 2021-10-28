package systemDepartment

import (
	"context"

	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemDepartmentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDepartmentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemDepartmentAddLogic {
	return SystemDepartmentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDepartmentAddLogic) SystemDepartmentAdd(req types.SystemDepartmentPostReq) (*types.SystemDepartmentReply, error) {
	// todo: add your logic here and delete this line

	return &types.SystemDepartmentReply{}, nil
}
