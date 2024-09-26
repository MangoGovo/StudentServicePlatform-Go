package service

import "StuService-Go/internal/model"

func GetUserByUserName(username string) (*model.User, error) {
	return d.GetUserByUserName(ctx, username)
}

func GetUserByID(UserID int64) (*model.User, error) {
	return d.GetUserByID(ctx, UserID)
}
