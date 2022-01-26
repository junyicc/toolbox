package db

import (
	"toolbox/logging"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql(mysqlConfig *MysqlConfig) *gorm.DB {
	if mysqlConfig.DBName == "" {
		return nil
	}
	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), GormConfig())
	if err != nil {
		logging.Panic("open mysql error", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	return db
}
