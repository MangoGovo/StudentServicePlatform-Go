package service

import (
	"StuService-Go/internal/model"
)

func GetUserByUserName(username string) (*model.User, error) {
	return d.GetUserByUserName(ctx, username)
}

func GetUserByID(UserID int64) (*model.User, error) {
	return d.GetUserByID(ctx, UserID)
}

func Register(Username string, NickName string, Password string, UserType int, AdminPwd string) error {

	// 检查账号密码格式
	// 1. 账号只能数字
	//for _, char := range Username {
	//	if char < '0' || char > '9' {
	//		utils.Log.Printf("%s注册请求不合法\n", Username)
	//		return &invalidUsernameOrPasswordError{}
	//	}
	//}
	//
	//// 2. 密码(md5)32位，数字加小写字母
	//if len(Password) != 32 {
	//	utils.Log.Printf("%s注册请求不合法\n", Username)
	//	return &invalidUsernameOrPasswordError{}
	//}
	//for _, char := range Password {
	//	if !(char >= '0' && char <= '9' || char >= 'a' && char <= 'z') {
	//		utils.Log.Printf("%s注册请求不合法\n", Username)
	//		return &invalidUsernameOrPasswordError{}
	//	}
	//}
	//
	//// 3. 检查usertype
	//if UserType != 1 && UserType != 2 {
	//	utils.Log.Printf("%s注册请求不合法\n", Username)
	//	return &invalidUserTypeError{}
	//}
	//utils.Log.Printf("%s注册成功，噢耶！\n", Username)

	return d.CreateUser(ctx, &model.User{
		Username: Username,
		Nickname: NickName,
		Password: Password,
		UserType: UserType,
	})
}
