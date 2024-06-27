package shopCar

import (
	"context"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShopCarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShopCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShopCarLogic {
	return &DeleteShopCarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteShopCarLogic) DeleteShopCar(req *types.GoodsInfo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	userId := uid.(int64)
	_, err = l.svcCtx.ShopCarRpc.DeleteGoodsToCar(l.ctx, &shopCar.ShopCarItemInfo{Goods: req.GoodsId, User: userId})
	if err != nil {
		return nil, err
	}
	return &types.Response{Msg: "删除成功"}, nil
}
