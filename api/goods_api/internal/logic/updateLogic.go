package logic

import (
	"context"
	"errors"
	"zero-mall/api/goods_api/internal/svc"
	"zero-mall/api/goods_api/internal/types"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	role := l.ctx.Value("role")
	if role != 2 {
		return nil, errors.New("权限不足")
	}
	_, err = l.svcCtx.GoodsRpc.UpdateGoods(l.ctx, &goods.CreateGoodsInfo{
		Id:         req.Id,
		Name:       req.Name,
		GoodsSn:    req.GoodsSn,
		GoodsBrief: req.GoodsBrief,
		ShopPrice:  req.ShopPrice,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
