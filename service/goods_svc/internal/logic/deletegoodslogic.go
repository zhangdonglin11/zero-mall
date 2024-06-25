package logic

import (
	"context"
	"errors"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsLogic) DeleteGoods(in *goods.DeleteGoodsInfo) (*goods.Empty, error) {
	// todo: add your logic here and delete this line
	if result := l.svcCtx.DB.Delete(&model.GoodsInfo{}, in.Id); result.RowsAffected == 0 {
		return nil, errors.New("商品不存在")
	}
	return &goods.Empty{}, nil
}
