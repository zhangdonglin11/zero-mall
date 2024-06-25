package logic

import (
	"context"
	"errors"
	"zero-mall/service/inventory_svc/internal/svc"
	"zero-mall/service/inventory_svc/model"
	"zero-mall/service/inventory_svc/types/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvDetailLogic {
	return &InvDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvDetailLogic) InvDetail(in *inventory.GoodsIdInfo) (*inventory.GoodsIdInfo, error) {
	// todo: add your logic here and delete this line
	var incDetail model.InventoryInfo
	if result := l.svcCtx.DB.Where(&model.InventoryInfo{Goods: in.GoodsId}).First(&incDetail); result.RowsAffected == 0 {
		return nil, errors.New("商品库存不存在")
	}
	return &inventory.GoodsIdInfo{
		GoodsId: incDetail.Goods,
		Num:     int64(incDetail.Stock),
	}, nil
}
