package dao

import (
	//"StuService-Go/internal/handler"
	"StuService-Go/internal/model"
	"context"
	//"gorm.io/gorm"
)

func (d *Dao) FeedbackCount(ctx context.Context) (int64, error) {
	var feedbackAmount int64
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("is_rubbish <> ?", 2).Count(&feedbackAmount).Error
	return feedbackAmount, err
}

func (d *Dao) UserCount(ctx context.Context) (int64, error) {
	var userAmount int64
	err := d.orm.WithContext(ctx).Model(&model.User{}).Count(&userAmount).Error
	return userAmount, err
}

func (d *Dao) RatingCount(ctx context.Context) ([]int, error) {
	var ratingAmount [5]int
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("status = ?", 3).Select("FeedbackRate,count(*)").Group("FeedbackRate").Scan(&ratingAmount).Error
	return []int{int(ratingAmount[0]), int(ratingAmount[1]), int(ratingAmount[2]), int(ratingAmount[3]), int(ratingAmount[4])}, err
}

func (d *Dao) GetUserList(ctx context.Context, UserType, PageCapacity, Offset int) ([]model.UserShow, error) {
	var users []model.UserShow
	var err error
	if UserType != -1 {
		err = d.orm.WithContext(ctx).Model(&model.User{}).Where("user_type = ?", UserType).Offset(Offset).Limit(PageCapacity).Find(&users).Error
	} else {
		err = d.orm.WithContext(ctx).Model(&model.User{}).Find(&users).Offset(Offset).Limit(PageCapacity).Error
	}
	return users, err
}

func (d *Dao) GetRubList(ctx context.Context, PageCapacity, Offset int) ([]model.Feedback, error) {
	var rubs []model.Feedback
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("is_rubbish = ?", 1).Offset(Offset).Limit(PageCapacity).Find(&rubs).Error
	return rubs, err
}

func (d *Dao) DealRubbish(ctx context.Context, FeedbackID int, IsRubbish bool) error {
	if IsRubbish {
		return d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("id = ?", FeedbackID).Update("IsRubbish", 2).Error
	} else {
		return d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("id = ?", FeedbackID).Update("IsRubbish", 0).Error
	}
}
