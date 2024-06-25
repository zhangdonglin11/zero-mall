package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"zero-mall/service/inventory_svc/internal/config"
	"zero-mall/service/inventory_svc/model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库连接失败：", err)
	}
	db.AutoMigrate(&model.InventoryInfo{}, &model.SellDetailInfo{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
