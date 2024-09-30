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
