package service

import (
	// "StuService-Go/internal/dao"
	"StuService-Go/internal/model"
)

func CountFeedback() (int64, error) {
	return d.FeedbackCount(ctx)
}

func CountUser() (int64, error) {
	return d.UserCount(ctx)
}

func CountRating() ([]int, error) {
	return d.RatingCount(ctx)
}

func GetUserList(UserType, PageCapacity, Offset int) ([]model.UserShow, error) {
	return d.GetUserList(ctx, UserType, PageCapacity, Offset)
}

func NewUser(userName, nickName, password, introduction string, userType, gender int) error {
	return d.CreateUser(ctx, &model.User{
		Username:     userName,
		Nickname:     nickName,
		Password:     password,
		UserType:     userType,
		Gender:       gender,
		Introduction: introduction,
	})
}
func DeleteUserBySudo(userID int) error {
	return d.DeleteUser(ctx, &model.User{
		ID: int64(userID),
	})
}

func GetRubList(PageCapacity, Offset int) ([]model.Feedback, error) {
	return d.GetRubList(ctx, PageCapacity, Offset)
}

func DealRubbish(FeedbackID int, IsRubbish bool) error {
	return d.DealRubbish(ctx, FeedbackID, IsRubbish)
}
