package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"reflect"
	"zero-mall/api/order_api/internal/config"
	"zero-mall/service/goods_svc/goodsclient"
	"zero-mall/service/inventory_svc/inventoryclient"
	"zero-mall/service/order_svc/orderclient"
	"zero-mall/service/shopcar_svr/shopcarclient"

	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ServiceContext struct {
	Config       config.Config
	ShopCarRpc   shopcarclient.ShopCar
	OrderRpc     orderclient.Order
	InventoryRpc inventoryclient.Inventory
	GoodsRpc     goodsclient.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ShopCarRpc:   shopcarclient.NewShopCar(zrpc.MustNewClient(c.ShopCarRpcConf)),
		GoodsRpc:     goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
		OrderRpc:     orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		InventoryRpc: inventoryclient.NewInventory(zrpc.MustNewClient(c.InventoryRpcConf)),
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
