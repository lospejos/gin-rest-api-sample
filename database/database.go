package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // configures postgres driver
	//	"github.com/velopert/gin-rest-api-sample/database/models"
	"gin-rest-api-sample/database/models"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	//db, err := gorm.Open("mysql", dbConfig)
	db, err := gorm.Open("postgres", dbConfig)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
