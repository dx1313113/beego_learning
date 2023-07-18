package initialize

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

func Log() *logs.BeeLogger {
	log := logs.NewLogger(10000)

	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":7,"color":true}`)

	runmode, _ := config.String("runmode")
	if runmode == "prod" {
		log.DelLogger("console")
	}
	return log
}
