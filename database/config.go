package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbGlobal *gorm.DB

func init() {
	dbGlobal = ConnectToDB()
}

func ConnectToDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port + " sslmode=" + sslmode +
		" TimeZone=" + timezone
	//dsn := "host=localhost user=dfs password=codehard dbname=go_crud port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	fmt.Println("data: ", dsn)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")
	return DB
}

func GetDBConnection() *gorm.DB {
	sqlDB, err := dbGlobal.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		dbGlobal = ConnectToDB()
		log.Fatal(err)
	}
	fmt.Println("NEWIN: ", dbGlobal)
	return dbGlobal
}
