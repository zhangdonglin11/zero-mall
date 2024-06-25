package model

import (
	"gorm.io/gorm"
)

type GoodsInfo struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50);not null;comment:商品名字"`
	GoodsSn    string `gorm:"type:varchar(50);not null;comment:商品唯一编号"`
	GoodsBrief string `gorm:"type:varchar(100);not null;comment:商品简要描述"`
	ShopPrice  int64  `gorm:"not null;comment:商品价格"`
	Status     uint   `gorm:"not null;comment:商品状态"`
}
