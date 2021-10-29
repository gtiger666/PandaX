package logic

import (
	"context"

	"Pandax/panda-rpc/resource/internal/svc"
	"Pandax/panda-rpc/resource/resource"

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

func (l *PingLogic) Ping(in *resource.Request) (*resource.Response, error) {
	// todo: add your logic here and delete this line

	return &resource.Response{}, nil
}
