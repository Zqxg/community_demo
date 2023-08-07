package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	TopicId int    //user表关联的外键id
	Status  int    `gorm:"default:0"` // 0 未完成 1 已完成
	Content string `gorm:"type:longtext"`
}
