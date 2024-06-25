package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"zero-mall/service/shopcar_svr/model"

	"zero-mall/service/shopcar_svr/internal/svc"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollBackDelShopCarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollBackDelShopCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollBackDelShopCarLogic {
	return &RollBackDelShopCarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回滚
func (l *RollBackDelShopCarLogic) RollBackDelShopCar(in *shopCar.ShopCarListRequest) (*shopCar.Empty, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DB.Begin()
	for _, item := range in.Data {
		shopCarInfo := &model.ShopCarInfo{
			Model: gorm.Model{
				ID: uint(item.Id),
			},
			Goods:   item.Goods,
			Nums:    item.Nums,
			User:    item.User,
			Checked: item.Checked,
		}

		if result := tx.Save(shopCarInfo); result.RowsAffected == 0 {
			tx.Callback()
			return nil, status.Errorf(codes.Internal, "购物车商品回滚失败！")
		}
	}
	tx.Commit()
	return &shopCar.Empty{}, nil
}
