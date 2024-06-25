package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"reflect"
	"zero-mall/api/goods_api/internal/config"
	"zero-mall/service/goods_svc/goodsclient"
	"zero-mall/service/goods_svc/types/goods"

	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goods.GoodsClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
	}
}

// 验证
func (svc ServiceContext) Validate(dataStruct interface{}) error {
	zh_ch := zh.New()
	validate := validator.New()
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	// 验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
