package global

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis/v8"
)

var (
	BeeDb   *sql.DB
	RedisDb *redis.Client
	BeeLog  *logs.BeeLogger
)
