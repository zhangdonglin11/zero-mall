package shopCar

import (
	"context"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShopCarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShopCarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShopCarLogic {
	return &UpdateShopCarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShopCarLogic) UpdateShopCar(req *types.GoodsInfo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//uid := l.ctx.Value("uid")
	//userId := uid.(int64)
	//carResp, err := l.svcCtx.OrderRpc.UpdateGoodsToCar(l.ctx, &order.ShopCarItemInfo{
	//	Goods:   req.GoodsId,
	//	User:    userId,
	//	Checked: req.Checked,
	//})
	//if err != nil {
	//	return nil, err
	//}
	return &types.Response{}, nil
}
