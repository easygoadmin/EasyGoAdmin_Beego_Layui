package main

import (
	_ "easygoadmin/initialize/config"
	_ "easygoadmin/initialize/mysql"
	_ "easygoadmin/initialize/session"
	_ "easygoadmin/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true
	// 启动应用
	beego.Run()
}
