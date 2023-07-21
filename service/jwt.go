package service

import (
	"beego_learning/global"
	"context"
	"time"
)

var ctx = context.Background()

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.RedisDb.Get(ctx, userName).Result()
	return err, redisJWT
}

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.DEFAULT_EXPIRE_SECONDS) * time.Second
	err = global.RedisDb.Set(ctx, userName, jwt, timer).Err()
	return err
}
