package logic

import (
	"context"
	"errors"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(in *goods.CreateGoodsInfo) (*goods.Empty, error) {
	// todo: add your logic here and delete this line
	// 查询商品是否存在
	var newGoods model.GoodsInfo
	if result := l.svcCtx.DB.First(&newGoods, in.Id); result.RowsAffected == 0 {
		return nil, errors.New("商品不存在")
	}
	newGoods.Name = in.Name
	newGoods.GoodsBrief = in.GoodsBrief
	newGoods.ShopPrice = in.ShopPrice
	newGoods.Status = uint(in.Status)
	if result := l.svcCtx.DB.Save(&newGoods); result.Error != nil {
		return nil, result.Error
	}
	return &goods.Empty{}, nil
}
