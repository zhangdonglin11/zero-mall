package handler

import (
	"net/http"
	"zero-mall/api/goods_api/internal/logic"
	"zero-mall/api/goods_api/internal/svc"
	"zero-mall/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		err := svcCtx.Validate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
