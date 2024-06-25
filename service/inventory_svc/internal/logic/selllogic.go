package logic

import (
	"context"
	"errors"
	"fmt"
	"zero-mall/service/inventory_svc/internal/svc"
	"zero-mall/service/inventory_svc/model"
	"zero-mall/service/inventory_svc/types/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellLogic {
	return &SellLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SellLogic) Sell(in *inventory.SellInfo) (*inventory.Empty, error) {
	// todo: add your logic here and delete this line
	// 库存扣减记录
	var sellDetail model.SellDetailInfo
	if result := l.svcCtx.DB.Where(&model.SellDetailInfo{OrderSn: in.OrderSn}).First(&sellDetail); result.RowsAffected != 0 {
		return nil, errors.New("订单编号已存在")
	}

	tx := l.svcCtx.DB.Begin()

	//barrier防止空补偿、空悬挂等具体看dtm官网即可，别忘记加barrier表在当前库中，因为判断补偿与要执行的sql一起本地事务
	//barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	//
	//barrier.Call(tx.Statement.ConnPool.(*sql.Tx), func(tx1 *sql.Tx) error {
	//
	//	return nil
	//})
	//

	for _, info := range in.GoodsInfo {
		//扣减库存
		var inv model.InventoryInfo
		if result := tx.Where("goods = ?", info.GoodsId).First(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New(fmt.Sprintf("goodsID:%s,商品库存不存在", inv.Goods))
		}
		if inv.Stock < uint(info.Num) {
			tx.Rollback()
			return nil, errors.New(fmt.Sprintf("goodsID:%s,商品库存不足", inv.Goods))
		}
		inv.Stock -= uint(info.Num)
		tx.Save(&inv)

		//处理库存扣减记录
		sellDetail.Detail = append(sellDetail.Detail, model.GoodsDetail{
			Goods: info.GoodsId,
			Num:   info.Num,
		})
	}
	//保存库存扣减记录
	sellDetail.Status = 1
	sellDetail.OrderSn = in.OrderSn
	if result := tx.Save(&sellDetail); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("库存记录保存失败")
	}
	tx.Commit()
	return &inventory.Empty{}, nil
}
