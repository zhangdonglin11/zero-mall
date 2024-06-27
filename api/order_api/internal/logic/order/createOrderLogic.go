package order

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"
	"zero-mall/service/order_svc/types/order"
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

	orderInfo, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.OrderInfoRequest{
		UserId:  int64(userId),
		Address: req.Address,
		Name:    req.Name,
		Mobile:  req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	// 支付逻辑
	// ·····

	return &types.CreateResponse{
		Id:      int(orderInfo.Id),
		PlayUrl: "http://play.xxxxx.xxxx/",
	}, nil
}
