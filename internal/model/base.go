package model

type BaseModel struct {
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"createdAt,omitempty"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
	DeletedAt int64 `json:"deletedAt,omitempty"`
}

//// BeforeCreate
//// 创建前自动填充创建时间和更新时间
//func (base *BaseModel) BeforeCreate(_ *gorm.DB) error {
//	now := utils.GetCurrentMs()
//	base.CreatedAt = now
//	base.UpdatedAt = now
//	return nil
//}

//// BeforeUpdate
//// 更新前自动填充更新时间
//func (base *BaseModel) BeforeUpdate(_ *gorm.DB) error {
//	base.UpdatedAt = utils.GetCurrentMs()
//	return nil
//}
