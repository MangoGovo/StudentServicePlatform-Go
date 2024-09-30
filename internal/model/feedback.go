package model

import "time"

type Reply struct {
	Sender   User   `json:"sender"`
	Receiver User   `json:"receiver"`
	Content  string `json:"content"`
}

type Feedback struct {
	ID              int64     `json:"feedback_id"`
	Sender          int64     `json:"sender"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	FeedbackTitle   string    `json:"feedback_title"`
	FeedbackType    int       `json:"feedback_type"`
	FeedbackContent string    `json:"feedback_content"`
	FeedbackRate    int       `json:"feedback_rate"`
	FeedbackReply   string    `json:"feedback_reply"` // []Reply
	Pictures        string    `json:"pictures"`       // []string
	IsEmergency     bool      `json:"is_emergency"`
	IsAnonymous     bool      `json:"is_anonymous"`
	Status          int       `json:"status" gorm:"default:0"` //0无人接单 1有管理员接单，暂未回复 2管理员已回复 3用户已评价(已解决)
}
