package model

import (
	"gorm.io/gorm"
)

// 评价
type Rating struct {
	gorm.Model
	Star int
}

// 统计
type Statistics struct {
	FeedbackAmount int   `json:"feedback_amount"`
	UserAmount     int   `json:"user_amount"`
	Ratings        []int `json:"ratings"`
}

type UserShow struct {
	ID       int64  `json:"user_id" gorm:"primary_key"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`                   //昵称
	UserType int    `json:"user_type" gorm:"default:0"` //用户类型
}
