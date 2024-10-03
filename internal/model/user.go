package model

import "gorm.io/gorm"

type User struct {
	DeletedAt    gorm.DeletedAt `json:"-"`
	ID           int64          `json:"user_id" gorm:"primary_key"`
	Username     string         `json:"username"`
	Nickname     string         `json:"nickname"`                   //昵称
	Password     string         `json:"-"`                          //密码
	Gender       int            `json:"gender" gorm:"default:0"`    //性别
	Introduction string         `json:"introduction"`               //简介
	UserType     int            `json:"user_type" gorm:"default:0"` //用户类型
}
