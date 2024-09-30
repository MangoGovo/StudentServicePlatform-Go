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
