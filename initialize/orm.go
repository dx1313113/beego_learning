package initialize

import (
	"database/sql"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
)

func OrmMysql() *sql.DB {
	//注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//数据库连接
	user, _ := config.String("mysqlUser")
	pwd, _ := config.String("mysqlPwd")
	host, _ := config.String("mysqlHost")
	port, _ := config.String("mysqlPort")
	dbname, _ := config.String("mysqlDb")

	//dbConn := "root:123456@tcp(127.0.0.1:3306)/beego_learning?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"

	orm.RegisterDataBase("default", "mysql", dbConn)

	if db, err := orm.GetDB(); err == nil {
		logs.GetLogger("ORM").Println("数据库连接成功")
		return db
	} else {
		//an official log.Logger with prefix ORM
		logs.GetLogger("ORM").Println("数据库连接出错")
		return nil
	}

}
