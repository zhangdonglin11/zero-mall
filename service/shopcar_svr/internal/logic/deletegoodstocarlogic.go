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

type DeleteGoodsToCarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsToCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsToCarLogic {
	return &DeleteGoodsToCarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除购物车中的商品
func (l *DeleteGoodsToCarLogic) DeleteGoodsToCar(in *shopCar.ShopCarItemInfo) (*shopCar.Empty, error) {
	// todo: add your logic here and delete this line
	if result := l.svcCtx.DB.Where(&model.ShopCarInfo{Goods: in.Goods, User: in.User}).Delete(&model.ShopCarInfo{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	return &shopCar.Empty{}, nil
}
