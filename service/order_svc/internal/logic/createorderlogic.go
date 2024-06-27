package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
	"zero-mall/service/goods_svc/types/goods"
	"zero-mall/service/inventory_svc/types/inventory"
	"zero-mall/service/order_svc/model"
	"zero-mall/service/shopcar_svr/types/shopCar"

	"zero-mall/service/order_svc/internal/svc"
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

func (l *CreateOrderLogic) CreateOrder(in *order.OrderInfoRequest) (*order.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line
	orderInfo := model.OrderInfo{
		User:         in.UserId,
		OrderSn:      handelOrderSn(in.UserId),
		Status:       "PAYING",
		SingerMobile: in.Mobile,
		SignerName:   in.Name,
		Address:      in.Address,
	}

	// 获取购物车商品列表
	shopCarGoodslist, err := l.svcCtx.ShopCarRpc.GetShopCarList(l.ctx, &shopCar.ShopCarItemInfo{
		User:    in.UserId,
		Checked: true,
	})
	if err != nil {
		return nil, err
	}

	var goodsIdArr []int64
	goodsMap := make(map[int64]int64)
	for _, item := range shopCarGoodslist.Data {
		goodsIdArr = append(goodsIdArr, item.Goods)
		goodsMap[item.Goods] = item.Nums
	}

	// 查询最新商品信息价格等
	goodsList, err := l.svcCtx.GoodsRpc.BatchGetGoods(l.ctx, &goods.BatchGoodsIdInfo{Id: goodsIdArr})
	if err != nil {
		return nil, err
	}

	//统计订单总价
	var amount int64
	var goodsIdInfo []*inventory.GoodsIdInfo
	// 商品订单库存
	var orderGoodsInfo []*model.OrderGoodsInfo
	for _, item := range goodsList.Data {
		goodsIdInfo = append(goodsIdInfo, &inventory.GoodsIdInfo{
			GoodsId: item.Id,
			Num:     goodsMap[item.Id],
		})
		//	订单价格累加
		amount = goodsMap[item.Id] * item.ShopPrice
		orderGoodsInfo = append(orderGoodsInfo, &model.OrderGoodsInfo{
			Order:      orderInfo.ID,
			Goods:      item.Id,
			GoodsName:  item.Name,
			GoodsPrice: item.ShopPrice,
			Nums:       goodsMap[item.Id],
		})
	}
	// 创建订单
	orderInfo.Amount = amount
	tx := l.svcCtx.DB.Begin()
	if result := tx.Create(&orderInfo); result.RowsAffected == 0 {
		tx.Callback()
		return nil, status.Errorf(codes.Internal, "订单创建失败")
	}
	// 创建订单商品镜像
	if result := tx.Create(&orderGoodsInfo); result.RowsAffected == 0 {
		tx.Callback()
		return nil, status.Errorf(codes.Internal, "订单创建失败")
	}

	// 扣减库存
	_, err = l.svcCtx.Inventory.Sell(l.ctx, &inventory.SellInfo{
		OrderSn:   orderInfo.OrderSn,
		GoodsInfo: goodsIdInfo,
	})

	if err != nil {
		tx.Callback()
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	//删除购物车商品
	_, err = l.svcCtx.ShopCarRpc.BatchDelShopCar(l.ctx, &shopCar.ShopCarListRequest{
		Data: shopCarGoodslist.Data,
	})

	tx.Commit()
	return &order.OrderInfoResponse{
		OrderSn: orderInfo.OrderSn,
		Id:      int64(orderInfo.ID),
	}, nil
}

// 生成订单号
func handelOrderSn(userId int64) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.New(rand.NewSource(time.Now().Unix()))
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}
