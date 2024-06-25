package shopCar

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-mall/api/order_api/internal/logic/shopCar"
	"zero-mall/api/order_api/internal/svc"
	"zero-mall/api/order_api/internal/types"
)

func UpdateShopCarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		err := svcCtx.Validate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shopCar.NewUpdateShopCarLogic(r.Context(), svcCtx)
		resp, err := l.UpdateShopCar(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
