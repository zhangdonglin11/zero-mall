package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"zero-mall/service/order_svc/model"

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
	var orderInfo model.OrderInfo
	if result := l.svcCtx.DB.First(&orderInfo, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	var orderGoods []model.OrderGoodsInfo
	l.svcCtx.DB.Where(model.OrderGoodsInfo{Order: orderInfo.ID}).Find(&orderGoods)
	var data []*order.OrderGoodsDetail
	for _, item := range orderGoods {
		data = append(data, &order.OrderGoodsDetail{
			Id:         int64(item.ID),
			OrderId:    item.OrderId,
			GoodsId:    item.Goods,
			GoodsName:  item.GoodsName,
			GoodsPrice: item.GoodsPrice,
			Nums:       item.Nums,
		})
	}

	return &order.OrderInfoDetailResponse{
		OrderInfo: &order.OrderInfoResponse{
			Id:      int64(orderInfo.ID),
			UserId:  orderInfo.User,
			OrderSn: orderInfo.OrderSn,
			PayType: orderInfo.PayType,
			Status:  orderInfo.Status,
			Total:   orderInfo.Amount,
			Address: orderInfo.Address,
			Name:    orderInfo.SignerName,
			Mobile:  orderInfo.SingerMobile,
			AddTime: strconv.FormatInt(orderInfo.CreatedAt.Unix(), 10),
		},
		Goods: data,
	}, nil
}
