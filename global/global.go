package global

import (
	"LibraryManagement/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	Db       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
)
