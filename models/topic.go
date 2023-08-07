package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model

	Title   string `gorm:"index;not null"`
	Status  int    `gorm:"default:0"` // 0 未完成 1 已完成
	Content string `gorm:"type:longtext"`
}
