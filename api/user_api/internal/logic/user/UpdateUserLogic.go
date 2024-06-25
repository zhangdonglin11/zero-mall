package user

import (
	"context"
	"time"
	"zero-mall/api/user_api/internal/svc"
	"zero-mall/api/user_api/internal/types"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserInfoRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	id := uid.(int32)

	//将前端传递过来的日期格式转换成int
	loc, _ := time.LoadLocation("Local") //local的L必须大写
	birthDay, _ := time.ParseInLocation("2006-01-02", req.Birthday, loc)

	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserInfo{
		Id:       id,
		NickName: req.Name,
		Gender:   req.Gender,
		BirthDay: uint64(birthDay.Unix()),
	})
	if err != nil {
		return "", err
	}
	return "修改成功", nil
}
