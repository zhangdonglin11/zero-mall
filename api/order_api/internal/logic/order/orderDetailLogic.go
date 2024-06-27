package order

import (
	"context"
	"zero-mall/service/order_svc/types/order"

	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailRequst) (resp *types.OrderInfoResponse, err error) {
	// todo: add your logic here and delete this line
	detail, err := l.svcCtx.OrderRpc.OrderDetail(l.ctx, &order.OrderInfoRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["order"] = detail.OrderInfo
	data["goods"] = detail.Goods

	return &types.OrderInfoResponse{
		Data: data,
	}, nil
}
