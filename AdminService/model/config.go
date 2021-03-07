package model

import "time"

type Model struct {
	ID       uint       `gorm:"column:id; primary_key" json:"id"`
	CreateAt time.Time  `gorm:"column:create_at" json:"createAt"`
	UpdateAt time.Time  `gorm:"column:update_at" json:"updateAt"`
	DeleteAt *time.Time `sql:"index" gorm:"column:deleted_at"   json:"deletedAt"`
}
