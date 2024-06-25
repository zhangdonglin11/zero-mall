package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"zero-mall/service/goods_svc/goodsclient"
	"zero-mall/service/shopcar_svr/internal/config"
	"zero-mall/service/shopcar_svr/model"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	GoodsRpc goodsclient.Goods
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
	db.AutoMigrate(&model.ShopCarInfo{})
	return &ServiceContext{
		Config:   c,
		DB:       db,
		GoodsRpc: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
	}
}
