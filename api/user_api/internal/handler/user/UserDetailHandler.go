package user

import (
	"net/http"
	"zero-mall/api/user_api/internal/logic/user"
	"zero-mall/api/user_api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
