package logic

import (
	"context"
	"errors"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsDetailLogic {
	return &GetGoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsDetailLogic) GetGoodsDetail(in *goods.GoodInfoRequest) (*goods.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line
	var newGoods model.GoodsInfo
	if result := l.svcCtx.DB.First(&newGoods, in.Id); result.RowsAffected == 0 {
		return nil, errors.New("商品不存在")
	}
	return &goods.GoodsInfoResponse{
		Id:         int64(newGoods.ID),
		Name:       newGoods.Name,
		GoodsSn:    newGoods.GoodsSn,
		GoodsBrief: newGoods.GoodsBrief,
		ShopPrice:  newGoods.ShopPrice,
		Status:     int64(newGoods.Status),
	}, nil
}
