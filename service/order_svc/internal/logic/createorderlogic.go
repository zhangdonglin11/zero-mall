package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/order_svc/internal/svc"
	"zero-mall/service/order_svc/model"
	"zero-mall/service/order_svc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单
func (l *CreateOrderLogic) CreateOrder(in *order.OrderInfoRequest) (*order.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DB.Begin()
	orderInfo := model.OrderInfo{
		User:         in.UserId,
		OrderSn:      in.OrderSn,
		Status:       "PAYING",
		Amount:       in.Amount,
		Address:      in.Address,
		SignerName:   in.Name,
		SingerMobile: in.Mobile,
	}
	if result := l.svcCtx.DB.Create(&orderInfo); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}

	var orderGoodsList []*model.OrderGoodsInfo
	for _, good := range in.OrderGoods {
		orderGoodsList = append(orderGoodsList, &model.OrderGoodsInfo{
			Order:      orderInfo.ID,
			Goods:      good.GoodsId,
			GoodsName:  good.GoodsName,
			GoodsPrice: good.GoodsPrice,
			Nums:       good.Nums,
		})
	}
	if result := tx.Create(&orderGoodsList); result.RowsAffected == 0 {
		tx.Callback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}
	tx.Commit()
	return &order.OrderInfoResponse{}, nil
}
