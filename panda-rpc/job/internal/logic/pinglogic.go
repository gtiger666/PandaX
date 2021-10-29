package logic

import (
	"context"

	"Pandax/panda-rpc/job/internal/svc"
	"Pandax/panda-rpc/job/job"

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

func (l *PingLogic) Ping(in *job.Request) (*job.Response, error) {
	// todo: add your logic here and delete this line

	return &job.Response{}, nil
}
