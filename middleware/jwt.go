package middleware

import (
	"errors"
	"github.com/beego/beego/v2/core/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT : header payload signature
// json web token: 标头 有效负载 签名
const (
	SecretKEY              string = "Dx1313113"
	DEFAULT_EXPIRE_SECONDS int    = 600 // 默认10分钟
	PasswordHashBytes             = 16
)

// MyCustomClaims
// This struct is the payload
// 此结构是有效负载
type MyCustomClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

// JwtPayload
// This struct is the parsing of token payload
// 此结构是对token有效负载的解析
type JwtPayload struct {
	Username  string `json:"username"`
	UserID    int    `json:"userID"`
	IssuedAt  int64  `json:"iat"` // 发布日期
	ExpiresAt int64  `json:"exp"` // 过期时间
}

func GenerateToken(userID int, expiredSeconds int) (tokenString string, err error) {
	// 如果没设置过期时间，默认为 DEFAULT_EXPIRE_SECONDS 600s
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	// 创建声明
	mySigningKey := []byte(SecretKEY)
	// 过期时间 = 当前时间（/s）+ expiredSeconds（/s）
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	logs.Info("Token 将到期于：", time.Unix(expireAt, 0))

	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    "dx",              // 发行者
			IssuedAt:  time.Now().Unix(), // 发布时间
			ExpiresAt: expireAt,          // 过期时间
		},
	}

	// 利用上面创建的声明 生成token
	// NewWithClaims(签名算法 SigningMethod, 声明 Claims) *Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 利用密钥对token签名
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "",
			errors.New("错误: token生成失败！")
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string) (*JwtPayload, error) {
	// 获取编码前的token信息
	token, err := jwt.ParseWithClaims(tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})
	// 获取payload-声明内容
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		logs.Info("%v %v",
			claims.UserID,
			claims.StandardClaims.ExpiresAt, // 过期时间
		)
		logs.Info("Token 将过期于：",
			time.Unix(claims.StandardClaims.ExpiresAt, 0),
		)
		return &JwtPayload{
			Username:  claims.StandardClaims.Issuer, // 用户名：发行者
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {
		logs.Info(err.Error())
		return nil, errors.New("错误: token验证失败")
	}
}

// RefreshToken
// @Title RefreshToken
// @Description "更新token"
// @Param tokenString 		string 		"编码后的token"
// @return   newTokenString string    "编码后的新的token"
// @return   err   			error     "错误信息"
func RefreshToken(tokenString string) (newTokenString string, err error) {
	// 获取上一个token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})
	// 获取上一个token 的 payload-声明
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}

	// 创建新的声明
	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix() //new expired
	newClaims := MyCustomClaims{
		claims.UserID,
		jwt.StandardClaims{
			Issuer:    claims.StandardClaims.Issuer, //name of token issue
			IssuedAt:  time.Now().Unix(),            //time of token issue
			ExpiresAt: expireAt,
		},
	}

	// 利用新的声明，生成新的token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	// 利用签名算法对新的token进行签名
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("错误: 新的新json web token 生成失败！")
	}

	return tokenStr, nil
}
