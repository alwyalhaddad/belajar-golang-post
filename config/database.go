package config

import (
	"time"

	"github.com/alwyalhaddad/belajar-golang-post/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dialect := mysql.Open("root:alwy1030@tcp(localhost:3306)/belajar_golang_post?charset=utf8mb4&parsetime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	utils.PanicIfError(err)

	// Connection pool
	sqlDB, err := db.DB()
	utils.PanicIfError(err)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
