package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"zero-mall/service/goods_svc/internal/config"
	"zero-mall/service/goods_svc/model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Rds    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("mysql数据库连接失败！")
	}
	db.AutoMigrate(&model.GoodsInfo{})

	return &ServiceContext{
		Config: c,
		DB:     db,
		Rds:    redis.MustNewRedis(c.RedisConf),
	}
}
