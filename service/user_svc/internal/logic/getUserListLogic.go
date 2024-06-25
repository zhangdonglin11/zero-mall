package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-mall/service/goods_svc/model"
	"zero-mall/service/user_svc/internal/svc"
	model2 "zero-mall/service/user_svc/model"
	"zero-mall/service/user_svc/types/user"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func handelUserInfoResponse(u model2.UserInfo) *user.UserInfoResponse {

	userinfo := user.UserInfoResponse{
		Id:       int32(u.ID),
		PassWord: u.Password,
		Mobile:   u.Mobile,
		NickName: u.Nickname,
		Gender:   u.Gender,
		Role:     int32(u.Role),
	}
	if u.Birthday != nil {
		userinfo.BirthDay = uint64(u.Birthday.Unix())
	}
	return &userinfo
}

func (l *GetUserListLogic) GetUserList(in *user.PageInfo) (*user.UserListResponse, error) {
	// todo: add your logic here and delete this line
	var users []model2.UserInfo
	//查询总数
	result := l.svcCtx.DB.Find(&model2.UserInfo{})
	if result.Error != nil {
		return nil, result.Error
	}
	var userList user.UserListResponse
	userList.Total = int32(result.RowsAffected)
	// 分页查询
	l.svcCtx.DB.Scopes(model.Paginate(int(in.Pn), int(in.PSize))).Find(&users)
	for _, u := range users {
		userList.Data = append(userList.Data, handelUserInfoResponse(u))
	}

	return &userList, nil
}
