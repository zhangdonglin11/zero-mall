package shopCar

import (
	"context"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopCarListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopCarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopCarListLogic {
	return &ShopCarListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopCarListLogic) ShopCarList(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//uid := l.ctx.Value("uid")
	//userId := uid.(int64)
	//shopCarList, err := l.svcCtx.OrderRpc.GetShopCarList(l.ctx, &order.ShopCarItemInfo{
	//	User: userId,
	//})
	//if err != nil {
	//	return nil, err
	//}

	return &types.Response{}, nil
}
