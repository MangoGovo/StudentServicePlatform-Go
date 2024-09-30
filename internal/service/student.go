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
