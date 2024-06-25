package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"zero-mall/api/goods_api/internal/svc"
	"zero-mall/api/goods_api/internal/types"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("+++++++++")
	roleValue, _ := l.ctx.Value("role").(json.Number)
	roleInt, _ := strconv.Atoi(string(roleValue))
	fmt.Println(roleInt)
	if roleInt != 2 {
		return nil, errors.New("权限不足")
	}

	respGoods, err := l.svcCtx.GoodsRpc.CreateGoods(l.ctx, &goods.CreateGoodsInfo{
		Name:       req.Name,
		GoodsSn:    req.GoodsSn,
		GoodsBrief: req.GoodsBrief,
		ShopPrice:  req.ShopPrice,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: respGoods.Id}, nil
}
