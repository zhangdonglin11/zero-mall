package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"zero-mall/service/user_svc/internal/config"
	"zero-mall/service/user_svc/model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("mysql数据库连接失败！")
	}
	db.AutoMigrate(&model.UserInfo{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
