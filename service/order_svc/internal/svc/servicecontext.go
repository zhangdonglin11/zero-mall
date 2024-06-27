package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"zero-mall/service/goods_svc/goodsclient"
	"zero-mall/service/inventory_svc/inventoryclient"
	"zero-mall/service/order_svc/internal/config"
	"zero-mall/service/order_svc/model"
	"zero-mall/service/shopcar_svr/shopcarclient"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	ShopCarRpc shopcarclient.ShopCar
	GoodsRpc   goodsclient.Goods
	Inventory  inventoryclient.Inventory
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&model.OrderInfo{}, &model.OrderGoodsInfo{})
	return &ServiceContext{
		Config:     c,
		DB:         db,
		ShopCarRpc: shopcarclient.NewShopCar(zrpc.MustNewClient(c.ShopCartConf)),
		GoodsRpc:   goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
		Inventory:  inventoryclient.NewInventory(zrpc.MustNewClient(c.InventoryConf)),
	}
}
