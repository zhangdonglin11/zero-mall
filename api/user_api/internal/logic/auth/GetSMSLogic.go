package auth

import (
	"context"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"regexp"
	"zero-mall/api/user_api/common/sms"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSMSLogic {
	return &GetSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSMSLogic) GetSMS(req *types.SmsRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	reg_tel := regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`)
	if !reg_tel.MatchString(req.Mobile) {
		return "", errors.New("电话号码格式错误")
	}
	timeOut, _ := l.svcCtx.RedisDb.TtlCtx(l.ctx, req.Mobile)
	// 验证码有效300秒
	if timeOut > 240 {
		return "", errors.New("一分钟内只能获取一次验证码")
	}

	//生成验证码
	code := sms.GenerateAnSMSCode(5)
	//创建短信服务
	client, _err := sms.CreateClient(l.svcCtx.Config.SMS.AccessKeyID, l.svcCtx.Config.SMS.AccessKeySecret)
	if _err != nil {
		return "", errors.New("短信服务错误")
	}
	params := sms.CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["PhoneNumbers"] = tea.String(req.Mobile)
	queries["SignName"] = tea.String("宠物萌小程序")
	queries["TemplateCode"] = tea.String("SMS_464376212")
	//fmt.Sprintf("{\"code\":\"%s\"}", code)
	queries["TemplateParam"] = tea.String("{\"code\":\"" + code + "\"}")
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return "", errors.New("短信服务错误")
	}
	err = l.svcCtx.RedisDb.SetexCtx(l.ctx, req.Mobile, code, 300)
	if err != nil {
		return "", errors.New("redis内部错误")
	}

	return "验证码发送成功", nil
}
