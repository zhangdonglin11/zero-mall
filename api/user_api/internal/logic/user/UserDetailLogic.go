package user

import (
	"context"
	"time"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail() (resp *types.UserDetailResponse, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	id := uid.(int32)
	userDetail, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &types.UserDetailResponse{
		Id:       int64(userDetail.Id),
		Name:     userDetail.NickName,
		Gender:   userDetail.Gender,
		Mobile:   userDetail.Mobile,
		Birthday: time.Unix(int64(userDetail.BirthDay), 0).Format("2006-01-02"),
	}, nil
}
