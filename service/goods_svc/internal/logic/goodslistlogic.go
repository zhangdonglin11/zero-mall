package logic

import (
	"context"
	"gorm.io/gorm"
	"zero-mall/service/goods_svc/internal/svc"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/goods_svc/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品接口
func (l *GoodsListLogic) GoodsList(in *goods.GoodsFilterRequest) (*goods.GoodsListResponse, error) {
	// todo: add your logic here and delete this line
	var goodsList []*model.GoodsInfo
	result := l.svcCtx.DB.Scopes(Paginate(int(in.Page), int(in.PageSize))).Find(&goodsList)
	goodsListResponse := goods.GoodsListResponse{}
	for _, good := range goodsList {
		goodsListResponse.Data = append(goodsListResponse.Data, &goods.GoodsInfoResponse{
			Id:         int64(good.ID),
			Name:       good.Name,
			GoodsSn:    good.GoodsSn,
			GoodsBrief: good.GoodsBrief,
			ShopPrice:  good.ShopPrice,
			Status:     int64(good.Status),
		})
	}
	goodsListResponse.Total = result.RowsAffected
	return &goodsListResponse, nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
