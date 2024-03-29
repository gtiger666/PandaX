// Code generated by goctl. DO NOT EDIT!
// Source: develop.proto

package server

import (
	"context"

	"Pandax/panda-rpc/develop/develop"
	"Pandax/panda-rpc/develop/internal/logic"
	"Pandax/panda-rpc/develop/internal/svc"
)

type DevelopServer struct {
	svcCtx *svc.ServiceContext
}

func NewDevelopServer(svcCtx *svc.ServiceContext) *DevelopServer {
	return &DevelopServer{
		svcCtx: svcCtx,
	}
}

func (s *DevelopServer) Ping(ctx context.Context, in *develop.Request) (*develop.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
