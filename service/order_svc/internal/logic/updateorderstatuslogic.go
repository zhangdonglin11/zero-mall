package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/order_svc/model"

	"zero-mall/service/order_svc/internal/svc"
	"zero-mall/service/order_svc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *order.OrderStatus) (*order.Empty, error) {
	// todo: add your logic here and delete this line
	var orderInfo model.OrderInfo
	if result := l.svcCtx.DB.First(&orderInfo, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	orderInfo.Status = in.Status
	l.svcCtx.DB.Save(&orderInfo)
	return &order.Empty{}, nil
}
