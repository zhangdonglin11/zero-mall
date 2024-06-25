package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	// 验证码对比
	code, err := l.svcCtx.RedisDb.GetCtx(l.ctx, req.Mobile)
	fmt.Println("==============")
	fmt.Println(code)
	fmt.Println(req.Mobile)
	fmt.Println(req.Code)

	if err != nil || code != req.Code {
		return "", errors.New("验证码错误")
	}
	_, err = l.svcCtx.UserRpc.GetUserByMobile(l.ctx, &user.MobileRequest{Mobile: req.Mobile})
	if err == nil {
		return "", errors.New("用户已注册")
	}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	_, err = l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserInfo{
		Mobile:   req.Mobile,
		PassWord: req.Password,
		NickName: fmt.Sprintf("用户%d", random.Intn(900000)+100000),
	})
	if err != nil {
		return "", errors.New("注册失败")
	}
	l.svcCtx.RedisDb.DelCtx(l.ctx, req.Mobile)
	return "注册成功", nil
}
