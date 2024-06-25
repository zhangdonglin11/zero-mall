package model

import "gorm.io/gorm"

type ShopCarInfo struct {
	gorm.Model
	User    int64 `gorm:"type:int;index"` //在购物车列表中我们需要查询当前用户的购物车记录
	Goods   int64 `gorm:"type:int;index"` //加索引：我们需要查询时候， 1. 会影响插入性能 2. 会占用磁盘
	Nums    int64 `gorm:"type:int"`
	Checked bool  //是否选中
}
