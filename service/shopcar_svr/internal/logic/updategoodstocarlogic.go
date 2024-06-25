package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/shopcar_svr/model"

	"zero-mall/service/shopcar_svr/internal/svc"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsToCarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsToCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsToCarLogic {
	return &UpdateGoodsToCarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改购物车商品数量和选中状态
func (l *UpdateGoodsToCarLogic) UpdateGoodsToCar(in *shopCar.ShopCarItemInfo) (*shopCar.Empty, error) {
	// todo: add your logic here and delete this line
	var shopCarInfo model.ShopCarInfo
	if result := l.svcCtx.DB.Where(&model.ShopCarInfo{User: in.User, Goods: in.Goods}).First(&shopCarInfo); result.RowsAffected != 1 {
		return nil, status.Errorf(codes.NotFound, "购物车商品不存在")
	}
	shopCarInfo.Nums = in.Nums
	shopCarInfo.Checked = in.Checked
	l.svcCtx.DB.Save(&shopCarInfo)
	return &shopCar.Empty{}, nil
}
