package service

import (
	"StuService-Go/internal/model"
)

func CreateFeedback(feedback *model.Feedback) error {
	return d.CreateFeedback(ctx, feedback)

}

func GetFeedbackByID(ID int64) (*model.Feedback, error) {
	return d.GetFeedbackByID(ctx, ID)

}

func DeleteFeedback(feedback *model.Feedback) error {
	return d.DeleteFeedback(ctx, feedback)
}

func UpdateFeedback(feedback *model.Feedback) error {
	return d.UpdateFeedback(ctx, feedback)

}

func GetFeedbackList(userID int64, status int, capacity int, offset int) ([]model.Feedback, error) {
	return d.GetFeedbackList(ctx, userID, status, capacity, offset)
}

func GetFeedbackCount(userID int64, status int) (int64, error) {
	return d.GetFeedbackCount(ctx, userID, status)
}

func GetCommentsByFeedbackID(feedbackID int64) ([]model.Comment, error) {
	return d.GetCommentByFeedbackID(ctx, feedbackID)
}

func CreateComment(comment *model.Comment) error {
	return d.CreateComment(ctx, comment)
}

func GetCommentByID(ID int64) (*model.Comment, error) {
	return d.GetCommentByID(ctx, ID)

}

func DeleteComment(comment *model.Comment) error {
	return d.DeleteComment(ctx, comment)
}

// GetCommentCount 获取一个帖子的评论数
func GetCommentCount(feedbackID int64) (int64, error) {
	return d.GetCommentsCount(ctx, feedbackID)
}
