package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-12306/common/xerr"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserByUsernameLogic {
	return &QueryUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserByUsernameLogic) QueryUserByUsername(in *pb.UserNameReq) (*pb.UserNameResp, error) {
	// todo: add your logic here and delete this line
	// 查询
	User, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if User != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.LOGIN_MAIL_NOT_EXIST), "用户明：%v不存在,err:%v,", err, in.Username)
	}
	var resp pb.UserNameResp
	_ = copier.Copy(&resp, &User)
	return &resp, nil
}
