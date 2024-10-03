package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         int64          `json:"comment_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	SenderID   int64          `json:"sender_id"`
	ReceiverID int64          `json:"receiver_id"`
	FeedbackID int64          `json:"feedback_id"`
	Content    string         `json:"content"`
}
