package database

import (
	"fmt"
	"log"
	"os"
	"project_article/internal/models"
	"time"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "github.com/lib/pq"
)

func ConnectToPostgresql() *gorm.DB {
	dbUSER := os.Getenv("DB_USER")
	dbPASWORD := os.Getenv("DB_PASSWORD")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbDBNAME := os.Getenv("DB_DBNAME")

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUSER, dbPASWORD, dbHOST, dbPORT, dbDBNAME)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // Output to standard logger
		logger.Config{
			SlowThreshold: time.Second, // Log queries slower than this
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color output
		},
	)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	fmt.Println("Error loading database : ", err.Error())
	// 	return nil
	// }
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    fmt.Println("Connected to PostgreSQL successfully!")

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{})


	return db
}
