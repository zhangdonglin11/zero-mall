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

type BatchDelShopCarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDelShopCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDelShopCarLogic {
	return &BatchDelShopCarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除购物车的商品
func (l *BatchDelShopCarLogic) BatchDelShopCar(in *shopCar.ShopCarListRequest) (*shopCar.Empty, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DB.Begin()
	for _, item := range in.Data {
		if result := tx.Where(&model.ShopCarInfo{Goods: item.Goods, User: item.User}).Delete(&model.ShopCarInfo{}); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "商品不存在！")
		}
	}
	tx.Commit()
	return &shopCar.Empty{}, nil
}
