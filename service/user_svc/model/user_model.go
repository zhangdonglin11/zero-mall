package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null;comment:'用户电话'"`
	Password string     `gorm:"type:varchar(100);not null;comment:'用户密码'"`
	Nickname string     `gorm:"type:varchar(20);comment:'用户昵称'"`
	Gender   string     `gorm:"type:varchar(6);default:'male';comment:'female表示女, male表示男'"`
	Birthday *time.Time `gorm:"type:datetime;"`
	Role     int        `gorm:"type:int;default:1;comment:'1表示普通用户, 2表示管理员'"`
}
