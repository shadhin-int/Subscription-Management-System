package migrations

import (
	_ "embed"
	"fmt"
	"log"
	db "subscription_management_system/database"
)

//go:embed query.sql
var loadData string

func LoadData() error {
	dbConnection := db.GetDBConnection()
	if dbConnection == nil {
		log.Fatal("Failed to connect database connection")
		return nil
	}
	dbC := dbConnection.Exec(loadData)
	if dbC.Error != nil {
		fmt.Println(dbC.Error.Error())
	}
	return nil
}
