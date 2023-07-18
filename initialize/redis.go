package initialize

import (
	"beego_learning/global"
	"context"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis/v8"
)

// 创建 redis 链接
func init() {
	addr, _ := config.String("redisAddr")
	pwd, _ := config.String("redisPwd")
	dbname, _ := config.Int("redisDb")

	var ctx = context.Background()
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,    // no password set
		DB:       dbname, // use default DB
	})
	_, err := global.RedisDb.Ping(ctx).Result()
	if err != nil {
		//连接失败
		println(err)
		logs.GetLogger("REDIS").Println("redis连接失败")
	} else {
		logs.GetLogger("REDIS").Println("redis连接成功")
	}
}
