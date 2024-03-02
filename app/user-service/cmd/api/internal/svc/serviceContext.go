package svc

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-12306/app/user-service/cmd/api/internal/config"
	"go-zero-12306/app/user-service/cmd/rpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(Token2uidInterceptor))),
	}
}
func Token2uidInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var uid string
	var MetadataKeyUid = "jwtUserId"
	ctxUid := ctx.Value(MetadataKeyUid)

	if ctxUid != nil {
		uid = ctxUid.(json.Number).String()
	}

	md := metadata.New(map[string]string{MetadataKeyUid: uid})

	ctx = metadata.NewOutgoingContext(ctx, md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	return nil
}
