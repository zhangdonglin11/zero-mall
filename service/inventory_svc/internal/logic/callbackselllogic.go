package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-mall/service/inventory_svc/internal/svc"
	"zero-mall/service/inventory_svc/model"
	"zero-mall/service/inventory_svc/types/inventory"
)

type CallbackSellLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackSellLogic {
	return &CallbackSellLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackSellLogic) CallbackSell(in *inventory.SellInfo) (*inventory.Empty, error) {
	// todo: add your logic here and delete this line
	//回滚扣减的库存
	//1.查询库存是否已回滚
	var stockSellDetail model.SellDetailInfo
	if result := l.svcCtx.DB.Where(&model.SellDetailInfo{OrderSn: in.OrderSn, Status: 1}).First(&stockSellDetail); result.RowsAffected == 0 {
		return nil, errors.New("没有需回滚库存记录")
	}
	tx := l.svcCtx.DB.Begin()
	//2.循环返回库存
	for _, goodsDetail := range stockSellDetail.Detail {
		var inv model.InventoryInfo
		if result := l.svcCtx.DB.Where(&model.InventoryInfo{Goods: goodsDetail.Goods}).First(&inv); result.RowsAffected == 0 {
			fmt.Printf("goodsId:%s,不存在l", goodsDetail.Goods)
			tx.Callback()
			return nil, errors.New("商品库存不存在")
		}
		inv.Stock += uint(goodsDetail.Num)
		if result := tx.Save(&inv); result.RowsAffected == 0 {
			tx.Callback()
			return nil, errors.New(fmt.Sprintf("goodsId:%s,回滚库存失败\n", goodsDetail.Goods))
		}
	}
	stockSellDetail.Status = 2
	if result := tx.Save(&stockSellDetail); result.RowsAffected == 0 {
		tx.Callback()
		return nil, errors.New("修改库存扣减记录失败")
	}
	tx.Commit()

	return &inventory.Empty{}, nil
}
