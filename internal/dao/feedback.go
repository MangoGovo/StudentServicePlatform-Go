package dao

import (
	"StuService-Go/internal/model"
	"context"
)

func (d *Dao) CreateFeedback(ctx context.Context, feedback *model.Feedback) error {
	return d.orm.WithContext(ctx).Create(feedback).Error
}

func (d *Dao) UpdateFeedback(ctx context.Context, feedback *model.Feedback) error {
	return d.orm.WithContext(ctx).Save(feedback).Error
}

func (d *Dao) DeleteFeedback(ctx context.Context, feedback *model.Feedback) error {
	return d.orm.WithContext(ctx).Delete(feedback).Error
}

func (d *Dao) GetFeedbackByID(ctx context.Context, ID int64) (*model.Feedback, error) {
	var feedback model.Feedback
	err := d.orm.WithContext(ctx).Where("id=?", ID).First(&feedback).Error
	return &feedback, err
}

func (d *Dao) GetFeedbackList(ctx context.Context, userID int64, status int, capacity int, offset int) ([]model.Feedback, error) {
	var feedbackList []model.Feedback
	if status == -1 {
		err := d.orm.WithContext(ctx).
			Model(&model.Feedback{}).
			Where("sender = ?", userID).
			Order("created_at desc").
			Limit(capacity).
			Offset(offset).
			Find(&feedbackList).Error
		return feedbackList, err
	}

	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("sender = ? AND status = ?", userID, status).Order("created_at desc").Limit(capacity).Offset(offset).Find(&feedbackList).Error

	return feedbackList, err
}

func (d *Dao) GetFeedbackCount(ctx context.Context, userID int64) (int64, error) {
	var total int64
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("sender = ?", userID).Count(&total).Error

	return total, err
}

func (d *Dao) AdminGetFeedbackList(ctx context.Context, status int, capacity int, offset int) ([]model.Feedback, error) {
	var feedbackList []model.Feedback
	if status == -1 {
		err := d.orm.WithContext(ctx).
			Model(&model.Feedback{}).
			Order("created_at desc").
			Limit(capacity).Offset(offset).
			Find(&feedbackList).Error
		return feedbackList, err
	}
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("status = ?", status).Order("created_at desc").Limit(capacity).Offset(offset).Find(&feedbackList).Error
	return feedbackList, err
}

func (d *Dao) AdminGetFeedbackCount(ctx context.Context, status int) (int, error) {
	var total int64
	if status == -1 {
		err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Count(&total).Error
		return int(total), err

	}
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("status = ?", status).Count(&total).Error
	return int(total), err
}
