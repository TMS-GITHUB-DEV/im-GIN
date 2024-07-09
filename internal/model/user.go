package model

import (
	"gorm.io/gorm"
	"im-GIN/internal/utils"
)

type User struct {
	Id        uint64 `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Pwd       string `json:"pwd,omitempty"`
	Sex       uint8  `json:"sex,omitempty"`
	Email     string `json:"email,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	BaseModel
}

// BeforeCreate
// 用户插入前生成雪花id
func (user *User) BeforeCreate(_ *gorm.DB) error {
	user.Id = utils.GetSnowflakeID()
	return nil
}
