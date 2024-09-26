package service

import (
	"StuService-Go/internal/dao"
	"context"
	"gorm.io/gorm"
)

var (
	ctx = context.Background()
	d   *dao.Dao
)

func Init(db *gorm.DB) {
	d = dao.New(db)
}
