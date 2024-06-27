package logic

import (
	"context"
	"strconv"
	"zero-mall/service/order_svc/model"

	"zero-mall/service/order_svc/internal/svc"
	"zero-mall/service/order_svc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *order.OrderFilterRequest) (*order.OrderListResponse, error) {
	// todo: add your logic here and delete this line
	var ordersInfo []model.OrderInfo
	result := l.svcCtx.DB.Where(model.OrderInfo{User: in.UserId}).Find(&ordersInfo)
	var data []*order.OrderInfoResponse
	for _, item := range ordersInfo {
		data = append(data, &order.OrderInfoResponse{
			Id:      int64(item.ID),
			UserId:  item.User,
			OrderSn: item.OrderSn,
			PayType: item.PayType,
			Status:  item.Status,
			Total:   item.Amount,
			Address: item.Address,
			Name:    item.SignerName,
			Mobile:  item.SingerMobile,
			AddTime: strconv.FormatInt(item.CreatedAt.Unix(), 10),
		})
	}

	return &order.OrderListResponse{
		Total: result.RowsAffected,
		Data:  data,
	}, nil
}
