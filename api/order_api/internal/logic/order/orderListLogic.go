package order

import (
	"context"
	"encoding/json"
	"strconv"
	"zero-mall/service/order_svc/types/order"

	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderInfoResponse, err error) {
	// todo: add your logic here and delete this line
	juid := l.ctx.Value("uid").(json.Number)
	uid, err := strconv.Atoi(string(juid))
	if err != nil {
		return nil, err
	}
	orderList, err := l.svcCtx.OrderRpc.OrderList(l.ctx, &order.OrderFilterRequest{
		UserId: int64(uid),
	})
	if err != nil {
		return nil, err
	}
	return &types.OrderInfoResponse{
		Data: orderList,
	}, nil
}
