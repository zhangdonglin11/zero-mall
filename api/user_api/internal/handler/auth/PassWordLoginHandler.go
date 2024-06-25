package auth

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-mall/api/user_api/internal/logic/auth"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"
)

func PassWordLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 验证
		err := svcCtx.Validate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewPassWordLoginLogic(r.Context(), svcCtx)
		resp, err := l.PassWordLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
