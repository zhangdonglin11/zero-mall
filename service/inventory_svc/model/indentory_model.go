package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type InventoryInfo struct {
	gorm.Model
	Goods   int64 `gorm:"not null;index;unique"`
	Stock   uint  `gorm:"not null"`
	Version int   `gorm:"not null"` // 乐观锁
}

type SellDetailInfo struct {
	gorm.Model
	OrderSn string          `gorm:"not null;index;unique"` // 订单编号
	Status  int             `gorm:"not null;"`
	Detail  GoodsDetailList `gorm:"not null;"`
}

type GoodsDetail struct {
	Goods int64
	Num   int64
}

type GoodsDetailList []GoodsDetail

func (g GoodsDetailList) Value() (value driver.Value, err error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
