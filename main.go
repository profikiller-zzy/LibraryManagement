package main

import (
	"LibraryManagement/core"
	"LibraryManagement/global"
)

func main() {
	// 读取配置文件，并将配置文件写入全局变量
	global.Config = core.InitConfig()
	// 初始化日志，并将日志写入全局变量
	global.Log = core.InitLogger()
	// 连接mysql数据库，并将数据库写入全局变量
	global.Db = core.InitGorm()
	global.Log.Info("你好！")
}
