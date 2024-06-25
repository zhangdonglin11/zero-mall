package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var store = base64Captcha.DefaultMemStore

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaResponse, err error) {
	// todo: add your logic here and delete this line
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, code, err := cp.Generate()
	if err != nil {
		//zap.S().Errorf("生成验证码错误,: ", err.Error())
		return nil, errors.New("生成验证码错误")
	}
	fmt.Println(code)
	return &types.CaptchaResponse{
		CaptchaId: id,
		PicPath:   b64s,
	}, nil
}
