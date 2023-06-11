package models

type TimeStamp struct {
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}