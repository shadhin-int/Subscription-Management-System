package migrations

import (
	"database/sql"
	db "github.com/shadhin-int/Subscription-Management-System.git/database"
	"github.com/shadhin-int/Subscription-Management-System.git/models"
	"log"
)

func DoMigrate() error {
	customerModel := &models.Customer{}
	subscriptionModel := &models.Subscription{}
	contractModel := &models.Contract{}
	invoiceModel := &models.Invoice{}

	dbConnection := db.GetDBConnection()
	if dbConnection == nil {
		log.Fatal("Failed to connect database connection")
		return nil
	}

	err := dbConnection.AutoMigrate(customerModel, subscriptionModel, contractModel, invoiceModel)

	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Setting table options...")
	if err := dbConnection.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		customerModel, subscriptionModel, contractModel, invoiceModel,
	); err != nil {
		log.Printf("Error setting table options: %v\n", err)
		return nil
	}
	log.Println("Table options set successfully.")

	log.Println("Database migration completed successfully.")

	success, err := dbConnection.DB()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(success *sql.DB) {
		err := success.Close()
		if err != nil {

		}
	}(success)
	return nil
}
