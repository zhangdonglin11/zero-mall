package logic

import (
	"context"
	"errors"
	"zero-mall/api/goods_api/internal/svc"
	"zero-mall/api/goods_api/internal/types"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	// todo: add your logic here and delete this line
	role := l.ctx.Value("role")
	if role != 2 {
		return nil, errors.New("权限不足")
	}
	_, err = l.svcCtx.GoodsRpc.DeleteGoods(l.ctx, &goods.DeleteGoodsInfo{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
