// Code generated by goctl. DO NOT EDIT.
package types

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}

type LoginRequest struct {
	Mobile    string `json:"mobile" validate:"required,len=11"`
	Password  string `json:"password" validate:"required,min=6,max=16"`
	Captcha   string `json:"captcha" validate:"required,min=5,max=5"`
	CaptchaId string `json:"captcha_id" validate:"required"`
}

type RegisterRequest struct {
	Mobile   string `json:"mobile" validate:"required,len=11"`
	Password string `json:"password" validate:"required,min=6,max=16"`
	Code     string `json:"code" validate:"required"`
}

type SmsRequest struct {
	Mobile string `json:"mobile" validate:"required,len=11"`
	Type   uint   `json:"type" validate:"required"`
}

type UpdateUserInfoRequest struct {
	Name     string `json:"name" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Birthday string `json:"birthday" validate:"required"`
}

type UserDetailResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
	Birthday string `json:"birthday"`
}
