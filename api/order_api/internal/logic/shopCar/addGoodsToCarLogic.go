package shopCar

import (
	"context"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsToCarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodsToCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsToCarLogic {
	return &AddGoodsToCarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodsToCarLogic) AddGoodsToCar(req *types.GoodsInfo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	userId := uid.(int64)
	_, err = l.svcCtx.ShopCarRpc.AddGoodsToCar(l.ctx, &shopCar.ShopCarItemInfo{
		User:  userId,
		Goods: req.GoodsId,
		Nums:  req.Num,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Msg: "请求成功",
	}, nil
}
