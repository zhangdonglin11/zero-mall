package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mall/service/user_svc/common"
	"zero-mall/service/user_svc/internal/svc"
	"zero-mall/service/user_svc/model"
	"zero-mall/service/user_svc/types/user"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserInfo) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	//新建用户
	var newUser model.UserInfo
	result := l.svcCtx.DB.Where(&model.UserInfo{Mobile: in.Mobile}).First(&newUser)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	newUser.Mobile = in.Mobile
	newUser.Nickname = in.NickName

	//密码加密
	newUser.Password = common.PasswordEncrypt(l.svcCtx.Config.Salt, in.PassWord)

	result = l.svcCtx.DB.Create(&newUser)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return &user.UserInfoResponse{
		Id:       int32(newUser.ID),
		PassWord: newUser.Password,
		Mobile:   newUser.Mobile,
		NickName: newUser.Nickname,
		Gender:   newUser.Gender,
		Role:     int32(newUser.Role),
	}, nil

}
