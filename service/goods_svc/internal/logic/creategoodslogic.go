package logic

import (
	"context"
	"errors"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsLogic {
	return &CreateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsLogic) CreateGoods(in *goods.CreateGoodsInfo) (*goods.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line
	newGoods := model.GoodsInfo{
		Name:       in.Name,
		GoodsSn:    in.GoodsSn,
		GoodsBrief: in.GoodsBrief,
		ShopPrice:  in.ShopPrice,
		Status:     uint(in.Status),
	}
	if result := l.svcCtx.DB.Debug().Create(&newGoods); result.RowsAffected == 0 {
		return nil, errors.New("商品创建失败")
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
