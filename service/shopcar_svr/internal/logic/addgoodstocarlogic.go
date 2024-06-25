package logic

import (
	"context"
	"zero-mall/service/shopcar_svr/model"

	"zero-mall/service/shopcar_svr/internal/svc"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsToCarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsToCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsToCarLogic {
	return &AddGoodsToCarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商品到购物车
func (l *AddGoodsToCarLogic) AddGoodsToCar(in *shopCar.ShopCarItemInfo) (*shopCar.ShopCarItemInfo, error) {
	// todo: add your logic here and delete this line
	var goodsCarInfo shopCar.ShopCarItemInfo
	if result := l.svcCtx.DB.Where(model.ShopCarInfo{Goods: in.Goods, User: in.User}).First(&goodsCarInfo); result.RowsAffected != 0 {
		goodsCarInfo.Nums += in.Nums
	} else {
		goodsCarInfo.Goods = in.Goods
		goodsCarInfo.User = in.User
		goodsCarInfo.Nums = in.Nums
		goodsCarInfo.Checked = false
	}
	l.svcCtx.DB.Save(&goodsCarInfo)
	return &shopCar.ShopCarItemInfo{}, nil
}
