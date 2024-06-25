package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/model"
	"zero-mall/service/user_svc/types/user"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserInfo) (*user.Empty, error) {
	// todo: add your logic here and delete this line
	//个人中心更新用户
	var newUser model.UserInfo
	if result := l.svcCtx.DB.First(&newUser, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	newUser.Nickname = in.NickName
	if in.BirthDay != 0 {
		birthDay := time.Unix(int64(in.BirthDay), 0)
		newUser.Birthday = &birthDay
	}
	newUser.Gender = in.Gender

	if result := l.svcCtx.DB.Save(&newUser); result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}
