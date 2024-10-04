package dao

import (
	"StuService-Go/internal/model"
	"context"
)

func (d *Dao) CreateComment(ctx context.Context, comment *model.Comment) error {
	return d.orm.WithContext(ctx).Create(comment).Error
}

func (d *Dao) UpdateComment(ctx context.Context, comment *model.Comment) error {
	return d.orm.WithContext(ctx).Save(comment).Error
}

func (d *Dao) DeleteComment(ctx context.Context, comment *model.Comment) error {
	return d.orm.WithContext(ctx).Delete(comment).Error
}

func (d *Dao) GetCommentByFeedbackID(ctx context.Context, feedbackID int64) ([]model.Comment, error) {
	var commentList []model.Comment
	err := d.orm.WithContext(ctx).Model(&model.Comment{}).Where("feedback_id = ?", feedbackID).Find(&commentList).Error
	return commentList, err
}

func (d *Dao) GetCommentByID(ctx context.Context, ID int64) (*model.Comment, error) {
	var comment model.Comment
	err := d.orm.WithContext(ctx).Where("id=?", ID).First(&comment).Error
	return &comment, err
}

// GetCommentsCount 获取一个帖子的评论数
func (d *Dao) GetCommentsCount(ctx context.Context, feedbackID int64) (int64, error) {
	var total int64
	err := d.orm.WithContext(ctx).Model(&model.Comment{}).Where("feedback_id = ?", feedbackID).Count(&total).Error
	return total, err
}
