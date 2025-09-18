package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func KamipaNewDatabase(cnf *ConfigApp) *gorm.DB {
	username := cnf.KamipaDB.Username
	password := cnf.KamipaDB.Password
	host := cnf.KamipaDB.Host
	port := cnf.KamipaDB.Port
	dbname := cnf.KamipaDB.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to Kamipa database:", err)
	}

	return db
}

func SimipaNewDatabase(cnf *ConfigApp) *gorm.DB {
	username := cnf.SimipaDB.Username
	password := cnf.SimipaDB.Password
	host := cnf.SimipaDB.Host
	port := cnf.SimipaDB.Port
	dbname := cnf.SimipaDB.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to Kamipa database:", err)
	}

	return db
}
