package logic

import (
	"context"
	"zero-mall/service/inventory_svc/internal/svc"
	"zero-mall/service/inventory_svc/model"
	"zero-mall/service/inventory_svc/types/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetInvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInvLogic {
	return &SetInvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetInvLogic) SetInv(in *inventory.GoodsIdInfo) (*inventory.Empty, error) {
	// todo: add your logic here and delete this line
	// 保存商品库存
	var data model.InventoryInfo
	l.svcCtx.DB.Where(&model.InventoryInfo{Goods: in.GoodsId}).First(&data)
	data.Stock = uint(in.Num)
	l.svcCtx.DB.Save(&data)
	return &inventory.Empty{}, nil
}
