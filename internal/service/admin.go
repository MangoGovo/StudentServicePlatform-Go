package service

import "StuService-Go/internal/model"

func AdminGetFeedbackList(status int, capacity int, offset int) ([]model.Feedback, error) {
	return d.AdminGetFeedbackList(ctx, status, capacity, offset)
}

func AdminGetFeedbackCount(status int) (int, error) {
	return d.AdminGetFeedbackCount(ctx, status)
}
