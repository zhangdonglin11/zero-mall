package logic

import (
	"context"

	"zero-mall/service/order_svc/internal/svc"
	"zero-mall/service/order_svc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDetailLogic) OrderDetail(in *order.OrderInfoRequest) (*order.OrderInfoDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrderInfoDetailResponse{}, nil
}
