package repositoryImpl

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormConfig = &gorm.Config{
	Logger:         logger.Default.LogMode(logger.Silent),
	TranslateError: true,
}

func InitSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), gormConfig)
	if err != nil {
		panic(err)
	}
	return db
}
