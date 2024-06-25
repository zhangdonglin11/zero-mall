package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/model"
	"zero-mall/service/user_svc/types/user"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *user.MobileRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var newUser model.UserInfo
	result := l.svcCtx.DB.Where(model.UserInfo{Mobile: in.Mobile}).First(&newUser)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	userInfo := user.UserInfoResponse{
		Id:       int32(newUser.ID),
		PassWord: newUser.Password,
		Mobile:   newUser.Mobile,
		NickName: newUser.Nickname,
		Gender:   newUser.Gender,
		Role:     int32(newUser.Role),
	}
	if newUser.Birthday != nil {
		userInfo.BirthDay = uint64(newUser.Birthday.Unix())
	}

	return &userInfo, nil
}
