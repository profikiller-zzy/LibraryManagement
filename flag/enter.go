package flag

import (
	"flag"
)

type Options struct {
	DB    bool // -db 自动迁移数据表
	Admin bool // -a 命令行创建管理员
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	dbFlag := flag.Bool("db", false, "auto migrate database")
	adminFlag := flag.Bool("a", false, "create admin")
	flag.Parse()
	var option = Options{
		DB:    *dbFlag,
		Admin: *adminFlag,
	}
	Execute(option)
}

func Execute(options Options) {
	if options.DB {
		MakeMigration()
	}

	if options.Admin {
		CreateAdmin()
	}
}
