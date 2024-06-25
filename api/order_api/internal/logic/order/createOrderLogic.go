package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"strconv"
	"time"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"
	"zero-mall/service/goods_svc/types/goods"
	"zero-mall/service/inventory_svc/types/inventory"
	"zero-mall/service/order_svc/types/order"
	"zero-mall/service/shopcar_svr/types/shopCar"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var dtmServer = "etcd://127.0.0.1:2379/dtmservice"

func (l *CreateOrderLogic) CreateOrder(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	jsonNumber := l.ctx.Value("uid").(json.Number)
	userId, _ := strconv.Atoi(string(jsonNumber))
	orderSn := handelOrderSn(int64(userId))

	// 获取购物车选中的商品信息
	shopCarList, err := l.svcCtx.ShopCarRpc.GetShopCarList(l.ctx, &shopCar.ShopCarItemInfo{
		User:    int64(userId),
		Checked: true,
	})
	if err != nil {
		return nil, err
	}

	var goodsIdArr []int64
	goodsMap := make(map[int64]int64)
	for _, goodsItem := range shopCarList.Data {
		goodsIdArr = append(goodsIdArr, goodsItem.Goods)
		goodsMap[goodsItem.Goods] = goodsItem.Nums
	}
	// 获取商品信息
	goodsList, err := l.svcCtx.GoodsRpc.BatchGetGoods(l.ctx, &goods.BatchGoodsIdInfo{
		Id: goodsIdArr,
	})
	if err != nil {
		return nil, err
	}
	// 订单总价格
	var amount int64
	// 订单商品信息
	var orderGoods []*order.OrderGoodsDetail
	// 扣除库存的信息
	var goodsIdInfo []*inventory.GoodsIdInfo
	for _, goodsItem := range goodsList.Data {
		amount += goodsItem.ShopPrice * goodsMap[goodsItem.Id]
		orderGoods = append(orderGoods, &order.OrderGoodsDetail{
			GoodsId:    goodsItem.Id,
			GoodsName:  goodsItem.Name,
			GoodsPrice: goodsItem.ShopPrice,
			Nums:       goodsMap[goodsItem.Id],
		})
		goodsIdInfo = append(goodsIdInfo, &inventory.GoodsIdInfo{
			GoodsId: goodsItem.Id,
			Num:     goodsMap[goodsItem.Id],
		})
	}

	fmt.Println("-----------------------------------------")
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("下单异常超时")
	}
	inventoryRpcBusiServer, err := l.svcCtx.Config.InventoryRpcConf.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("下单异常超时")
	}
	shopCarRpcBusiServer, err := l.svcCtx.Config.ShopCarRpcConf.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("下单异常超时")
	}

	createOrderReq := &order.OrderInfoRequest{
		OrderSn:    orderSn,
		UserId:     int64(userId),
		Address:    req.Address,
		Name:       req.Name,
		Mobile:     req.Mobile,
		Amount:     amount,
		OrderGoods: orderGoods,
	}
	inventoryReq := &inventory.SellInfo{
		OrderSn:   orderSn,
		GoodsInfo: goodsIdInfo,
	}
	shopCarReq := &shopCar.ShopCarListRequest{
		Data: shopCarList.Data,
	}

	//这里只举了saga例子，tcc等其他例子基本没啥区别具体可以看dtm官网

	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/order.order/CreateOrder", orderRpcBusiServer+"/order.order/CallbackOrder", createOrderReq).
		Add(inventoryRpcBusiServer+"/inventory.Inventory/CallbackSell", inventoryRpcBusiServer+"/inventory.Inventory/CallbackSell", inventoryReq).
		Add(shopCarRpcBusiServer+"/shopCar.ShopCar/BatchDelShopCar", shopCarRpcBusiServer+"/shopCar.ShopCar/RollBackDelShopCar", shopCarReq)

	err = saga.Submit()
	dtmimp.FatalIfError(err)
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}

	return &types.CreateResponse{
		Id:      1,
		PlayUrl: "http://play.xxxxx.xxxx/",
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
