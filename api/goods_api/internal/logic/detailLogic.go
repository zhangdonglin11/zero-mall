package logic

import (
	"context"
	"zero-mall/api/goods_api/internal/svc"
	"zero-mall/api/goods_api/internal/types"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	goodsDetail, err := l.svcCtx.GoodsRpc.GetGoodsDetail(l.ctx, &goods.GoodInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.DetailResponse{
		Id:         goodsDetail.Id,
		Name:       goodsDetail.Name,
		GoodsSn:    goodsDetail.GoodsSn,
		GoodsBrief: goodsDetail.GoodsBrief,
		ShopPrice:  goodsDetail.ShopPrice,
		Status:     goodsDetail.Status,
	}, nil
}
