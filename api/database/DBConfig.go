package database

import (
	"goLoginBoiler/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	conn_str := "loginBoiler:loginBoilerpw@tcp(database:3306)/loginBoiler?parseTime=true"
	log.Println("Attempting database connection....")

	db_conn, err := gorm.Open(mysql.Open(conn_str), &gorm.Config{})

	if err != nil {
		log.Fatalln("There was an error getting a connection to the Database")
	}

	db = db_conn
	log.Println("Connected to database....")

	migrateModels()
}

func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}

	return db
}

func migrateModels() {
	db.AutoMigrate(models.User{})
}
