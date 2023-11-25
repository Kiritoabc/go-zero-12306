// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/logic"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UserServer) CheckLogin(ctx context.Context, in *pb.CheckLoginReq) (*pb.LoginResp, error) {
	l := logic.NewCheckLoginLogic(ctx, s.svcCtx)
	return l.CheckLogin(in)
}

func (s *UserServer) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.VoidResp, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}
