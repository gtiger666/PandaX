package logic

import (
	"context"

	"Pandax/panda-rpc/system/internal/svc"
	"Pandax/panda-rpc/system/system"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *system.Request) (*system.Response, error) {
	// todo: add your logic here and delete this line

	return &system.Response{}, nil
}
