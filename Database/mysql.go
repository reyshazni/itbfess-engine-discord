package Database

import (
	"ITBFess/Model/Entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DatabaseConfig struct {
	Username string
	Password string
	Database string
	Url      string
	Port     int
}

func (conf DatabaseConfig) getDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Url, conf.Port, conf.Database)
	return dsn
}

var database *gorm.DB

func loadDatabase() {
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: os.Getenv("DB_DSN"),
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal("Error loading Database", err)
	}
	db.AutoMigrate(Entity.User{})
	db.AutoMigrate(Entity.Menfess{})
	database = db
	println("LOAD TEST DB")
}

func GetDatabase() *gorm.DB {
	if database == nil {
		loadDatabase()
	}
	return database
}
