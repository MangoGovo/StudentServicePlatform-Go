package database

import (
	"StuService-Go/internal/model"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(model.User{}, model.Feedback{}, model.Comment{})
}
