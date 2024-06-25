package auth

import (
	"net/http"
	"zero-mall/api/user_api/internal/logic/auth"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSMSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		err := svcCtx.Validate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewGetSMSLogic(r.Context(), svcCtx)
		resp, err := l.GetSMS(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
