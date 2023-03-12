package app

import (
	"log"
	"os"
	"time"

	"cek/helper"
	"cek/model/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	// "api_ski:mdnfjkt45@tcp(103.103.192.24:4000)/ski?parseTime=true"
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// RUN before_auto_migrate.sql
	helper.RunSQLFromFile(database, "app/database/before_auto_migrate.sql")

	err = database.AutoMigrate(

		// GOLANG
		&domain.Note{},
		&domain.User{},
		&domain.Session{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	// RUN after_auto_migrate.sql
	helper.RunSQLFromFile(database, "app/database/after_auto_migrate.sql")

	return database
}
