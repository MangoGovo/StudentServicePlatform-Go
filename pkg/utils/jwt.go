package utils

import (
	"StuService-Go/internal/apiException"
	"StuService-Go/internal/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int64) (string, error) {
	secretKey := global.Config.GetString("JWT.Secret")
	expireHour := global.Config.GetInt("JWT.ExpireHour")
	claims := User{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                            // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                            // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

func ParseJwt(tokenString string) (*User, error) {
	secretKey := global.Config.GetString("JWT.Secret")
	token, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		Log.Println("jwt解析失败")
		Log.Println(err.Error())
		return nil, apiException.ServerError
	}
	if claims, ok := token.Claims.(*User); ok && token.Valid {
		return claims, nil
	} else {
		return nil, apiException.ServerError
	}
}
