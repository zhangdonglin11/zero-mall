package model

import (
	"gorm.io/gorm"
	"time"
)

type OrderInfo struct {
	gorm.Model
	User    int64  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"` //订单号，我们平台自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)'"`

	Status  string     `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo string     `gorm:"type:varchar(100) comment '交易号'"` //交易号就是支付宝的订单号
	Amount  int64      `gorm:"comment '订单金额'"`
	PayTime *time.Time `gorm:"type:datetime"`

	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SingerMobile string `gorm:"type:varchar(11)"`
}

type OrderGoodsInfo struct {
	gorm.Model
	Order uint  `gorm:"type:int;index"`
	Goods int64 `gorm:"type:int;index"`

	//把商品的信息保存下来了 ， 字段冗余， 高并发系统中我们一般都不会遵循三范式  做镜像 记录
	GoodsName  string `gorm:"type:varchar(100);index"`
	GoodsPrice int64
	Nums       int64 `gorm:"type:int"`
}
