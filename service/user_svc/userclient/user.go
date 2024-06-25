// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"
	"zero-mall/service/user_svc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckResponse     = user.CheckResponse
	CreateUserInfo    = user.CreateUserInfo
	Empty             = user.Empty
	IdRequest         = user.IdRequest
	MobileRequest     = user.MobileRequest
	PageInfo          = user.PageInfo
	PasswordCheckInfo = user.PasswordCheckInfo
	UpdateUserInfo    = user.UpdateUserInfo
	UserInfoResponse  = user.UserInfoResponse
	UserListResponse  = user.UserListResponse

	User interface {
		GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
		GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*Empty, error)
		CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultUser) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByMobile(ctx, in, opts...)
}

func (m *defaultUser) GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckPassWord(ctx, in, opts...)
}
