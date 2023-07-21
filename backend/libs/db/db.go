package db

import (
	"aegis_test/libs"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	log.Println("Trying to Connect to DB " + libs.DbName + " on " + libs.DbHost)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", libs.DbUser, libs.DbPassword, libs.DbHost, libs.DbPort, libs.DbName)
	db, dbError := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if dbError != nil {
		log.Println("Failed to Connect to DB Server!!!")
		log.Println(dbError)
		panic(dbError)
	}

	db.Debug()
	log.Println("Connected to DB DB " + libs.DbName + " on " + libs.DbHost + " successfully!")

	return db
}
