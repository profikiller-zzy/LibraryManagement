package main

import (
	"LibraryManagement/core"
	"LibraryManagement/flag"
	"LibraryManagement/global"
)

func main() {
	// 读取配置文件，并将配置文件写入全局变量
	global.Config = core.InitConfig()
	// 初始化日志，并将日志写入全局变量
	global.Log = core.InitLogger()
	// 连接mysql数据库，并将数据库写入全局变量
	global.Db = core.InitGorm()

	// 捕获命令行参数，并对不同命令行参数的值来执行不同的操作
	flag.Parse()
}
