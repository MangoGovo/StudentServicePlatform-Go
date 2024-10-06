package dao

import (
	"StuService-Go/internal/model"
	"context"
)

func (d *Dao) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username=?", username).First(&user).Error
	return &user, err
}

func (d *Dao) GetUserByID(ctx context.Context, ID int64) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id=?", ID).First(&user).Error
	if err != nil {
		return &model.User{
			ID:       0,
			Username: "用户已注销",
			Nickname: "用户已注销",
			UserType: 0,
		}, nil
	}
	return &user, err
}

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Create(user).Error
}

func (d *Dao) UpdateUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Save(user).Error
}

func (d *Dao) DeleteUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Delete(user).Error
}

func (d *Dao) GetStuStat(ctx context.Context, userID int64) (model.StuStat, error) {
	var stuStat model.StuStat
	//	1. sent
	err := d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("sender=?", userID).Count(&stuStat.Sent).Error
	if err != nil {
		return stuStat, err
	}
	//	2. ignored
	err = d.orm.WithContext(ctx).
		Model(&model.Feedback{}).
		Where("sender = ? AND is_rubbish = 2", userID).
		Count(&stuStat.Ignored).
		Error

	return stuStat, err
}

func (d *Dao) GetAdminStat(ctx context.Context, userID int64) (model.AdminStat, error) {
	var adminStat model.AdminStat

	//	1. handled 接单数
	err := d.orm.WithContext(ctx).
		Model(&model.Feedback{}).
		Where("handler=?", userID).
		Count(&adminStat.Handled).
		Error

	if err != nil {
		return adminStat, err
	}

	//	2. ratings
	ratingList := []int64{0, 0, 0, 0, 0}
	for i := 0; i < len(ratingList); i++ {
		var rate int64
		_ = d.orm.WithContext(ctx).
			Model(&model.Feedback{}).
			Where("handler = ? AND status = 3", userID).
			Where("feedback_rate = ?", i+1).Count(&rate).Error
		ratingList[i] = rate
	}
	adminStat.Ratings = ratingList
	return adminStat, err
}
