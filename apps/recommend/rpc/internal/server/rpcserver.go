// Code generated by goctl. DO NOT EDIT.
// Source: recommend.proto

package server

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/recommend/rpc/internal/logic"
	"github.com/zhoushuguang/lebron/apps/recommend/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/recommend/rpc/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) Ping(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
