package common

import "time"

type SqlModel struct {
	Id       int        `json:"id" gorm:"column:id;"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at;"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at;"`
}
