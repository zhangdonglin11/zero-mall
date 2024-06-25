package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/order_svc/model"

	"zero-mall/service/order_svc/internal/svc"
	"zero-mall/service/order_svc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackOrderLogic {
	return &CallbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackOrderLogic) CallbackOrder(in *order.OrderInfoRequest) (*order.Empty, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DB.Begin()
	var orderInfo model.OrderInfo
	if result := l.svcCtx.DB.Where(&model.OrderInfo{OrderSn: in.OrderSn, User: in.UserId}).First(&orderInfo); result.RowsAffected == 0 {
		return &order.Empty{}, status.Errorf(codes.NotFound, "订单不存在")
	}
	fmt.Println(orderInfo)
	tx.Where(&model.OrderGoodsInfo{Order: orderInfo.ID}).Delete(&model.OrderGoodsInfo{})
	if result := tx.Delete(&model.OrderInfo{}, orderInfo.ID); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "删除订单失败")
	}
	tx.Commit()
	return &order.Empty{}, nil
}
