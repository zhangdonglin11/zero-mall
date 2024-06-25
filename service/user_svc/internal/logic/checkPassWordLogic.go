package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-mall/service/user_svc/common"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/types/user"
)

type CheckPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPassWordLogic {
	return &CheckPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckPassWordLogic) CheckPassWord(in *user.PasswordCheckInfo) (*user.CheckResponse, error) {
	// todo: add your logic here and delete this line
	//密码校验
	encrypt := common.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

	return &user.CheckResponse{
		Success: in.EncryptedPassword == encrypt,
	}, nil
}
