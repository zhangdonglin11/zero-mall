package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"time"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassWordLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPassWordLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassWordLoginLogic {
	return &PassWordLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PassWordLoginLogic) PassWordLogin(req *types.LoginRequest) (string, error) {
	// todo: add your logic here and delete this line
	reg_tel := regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`)
	if !reg_tel.MatchString(req.Mobile) {
		return "", errors.New("电话号码格式错误")
	}
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		return "", errors.New("验证码错误")
	}
	// 查询用户是否存在
	if respUser, err := l.svcCtx.UserRpc.GetUserByMobile(l.ctx, &user.MobileRequest{
		Mobile: req.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return "", errors.New("用户不存在")
			default:
				return "", errors.New("登录失败")
			}
		}
	} else {
		// 判断密码
		if passResp, passErr := l.svcCtx.UserRpc.CheckPassWord(l.ctx, &user.PasswordCheckInfo{
			Password:          req.Password,
			EncryptedPassword: respUser.PassWord,
		}); passErr != nil {
			return "", errors.New("登录失败!")
		} else {
			if passResp.Success {
				//生成token
				claims := make(jwt.MapClaims)
				claims["exp"] = time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire
				claims["iat"] = time.Now().Unix()
				claims["uid"] = respUser.Id
				claims["role"] = respUser.Role
				token := jwt.New(jwt.SigningMethodHS256)
				token.Claims = claims
				return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
			}
		}

	}
	//生成token
	return "请求成功", nil
}
