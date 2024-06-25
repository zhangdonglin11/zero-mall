package logic

import (
	"context"
	"zero-mall/service/shopcar_svr/model"

	"zero-mall/service/shopcar_svr/internal/svc"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopCarListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopCarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopCarListLogic {
	return &GetShopCarListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取购物车中的商品 传用户id获取用户的购物车商品，+传checked获取购物车选中商品
func (l *GetShopCarListLogic) GetShopCarList(in *shopCar.ShopCarItemInfo) (*shopCar.ShopCarListResponse, error) {
	// todo: add your logic here and delete this line
	var shopCarInfoList []model.ShopCarInfo
	result := l.svcCtx.DB.Where(&model.ShopCarInfo{User: in.User, Checked: in.Checked}).Find(&shopCarInfoList)
	shopCarListResponse := &shopCar.ShopCarListResponse{
		Total: result.RowsAffected,
	}
	for _, item := range shopCarInfoList {
		shopCarListResponse.Data = append(shopCarListResponse.Data, &shopCar.ShopCarItemInfo{
			Id:      int64(item.ID),
			Goods:   item.Goods,
			User:    item.User,
			Nums:    item.Nums,
			Checked: item.Checked,
		})

	}
	return shopCarListResponse, nil
}
