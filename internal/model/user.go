package model

type User struct {
	ID           int64  `json:"user_id"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`                   //昵称
	Password     string `json:"-"`                          //密码
	Gender       int    `json:"gender" gorm:"default:0"`    //性别
	Introduction string `json:"introduction"`               //简介
	UserType     int    `json:"user_type" gorm:"default:0"` //用户类型
}
