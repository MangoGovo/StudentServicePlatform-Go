package service

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/dao"
	"StuService-Go/internal/model"
	"StuService-Go/pkg/utils"
)

func GetUserByUserName(username string) (*model.User, error) {
	return d.GetUserByUserName(ctx, username)
}

func GetUserByID(UserID int64) (*model.User, error) {
	return d.GetUserByID(ctx, UserID)
}

func Register(Username string, NickName string, Password string, UserType int) error {
	// 1. 检查邮箱格式
	if !utils.IsValidEmail(Username) {
		return apiException.ParamsError
	}

	// 2. 检查密码强弱
	// TODO 检查密码强弱

	// 3. 将密码hash后添加到数据库
	EncryptedPwd := utils.MD5(Password)

	// 4. 检查usertype
	if UserType != 0 {
		return apiException.PermissionsNotAllowed
	}

	// 5. 从Redis删除验证码
	if err := dao.RedisDelKeyVal(ctx, Username); err != nil {
		utils.Log.Println(err.Error())
		return apiException.ServerError
	}

	return d.CreateUser(ctx, &model.User{
		Username: Username,
		Nickname: NickName,
		Password: EncryptedPwd,
		UserType: UserType,
	})
}

func UpdateUser(user *model.User) error {
	return d.UpdateUser(ctx, user)
}
