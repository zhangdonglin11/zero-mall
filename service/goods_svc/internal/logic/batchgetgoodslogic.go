package logic

import (
	"context"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetGoodsLogic {
	return &BatchGetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品接口
func (l *BatchGetGoodsLogic) BatchGetGoods(in *goods.BatchGoodsIdInfo) (*goods.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	var goodsLsat []model.GoodsInfo
	result := l.svcCtx.DB.Where("id in (?)", in.Id).Find(&goodsLsat)

	resp := goods.GoodsListResponse{}
	for _, good := range goodsLsat {
		resp.Data = append(resp.Data, &goods.GoodsInfoResponse{
			Id:         int64(good.ID),
			Name:       good.Name,
			GoodsSn:    good.GoodsSn,
			GoodsBrief: good.GoodsBrief,
			ShopPrice:  good.ShopPrice,
			Status:     int64(good.Status),
		})
	}
	resp.Total = result.RowsAffected

	return &resp, nil
}
