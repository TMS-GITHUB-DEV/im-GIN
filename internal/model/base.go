package model

import (
	"gorm.io/gorm"
	"im-GIN/internal/utils"
)

type BaseModel struct {
	CreatedAt uint64 `json:"created_at,omitempty"`
	UpdatedAt uint64 `json:"updated_at,omitempty"`
	DeletedAt uint64 `json:"deleted_at,omitempty"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	now := utils.GetCurrentTs()
	base.CreatedAt = now
	base.UpdatedAt = now
	return nil
}

func (base *BaseModel) BeforeUpdate(_ *gorm.DB) error {
	base.UpdatedAt = utils.GetCurrentTs()
	return nil
}

//func (base *BaseModel) BeforeDelete(tx *gorm.DB) error {
//	base.DeletedAt = utils.GetCurrentTs()
//	return nil
//}
