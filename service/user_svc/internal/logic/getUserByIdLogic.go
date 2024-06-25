package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/model"
	"zero-mall/service/user_svc/types/user"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.IdRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var result model.UserInfo
	res := l.svcCtx.DB.First(&result, in.Id)
	if res.Error != nil {
		return nil, res.Error
	}
	userInfo := user.UserInfoResponse{
		Id:       int32(result.ID),
		PassWord: result.Password,
		Mobile:   result.Mobile,
		NickName: result.Nickname,
		Gender:   result.Gender,
		Role:     int32(result.Role),
	}
	if result.Birthday != nil {
		userInfo.BirthDay = uint64(result.Birthday.Unix())
	}
	return &userInfo, nil
}
