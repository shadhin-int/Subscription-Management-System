package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/robfig/cron"
	"gorm.io/gorm"
	"strconv"
	db "subscription_management_system/database"
	"subscription_management_system/generate_pdf"
	"subscription_management_system/models"
	"time"
)

func Customer() chi.Router {
	route := chi.NewRouter()

	return route
}

func Invoice() chi.Router {
	route := chi.NewRouter()

	return route
}

var dbCon *gorm.DB

func init() {
	dbCon = db.GetDBConnection()
}

func getNextBillingDate(contractStartDate time.Time) time.Time {
	year, month, _ := contractStartDate.Date()
	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		year++
	}
	nextMonthDays := time.Date(year, nextMonth, 0, 0, 0, 0, 0, time.UTC).Day()
	nextBillingDate := time.Date(year, nextMonth, contractStartDate.Day(), 0, 0, 0, 0, time.UTC)

	if contractStartDate.Day() > nextMonthDays {
		nextBillingDate = time.Date(year, nextMonth, nextMonthDays, 0, 0, 0, 0, time.UTC)
	}

	return nextBillingDate
}

func isBillingDue(nextBillingDate time.Time) bool {
	currentDate := time.Now()
	return currentDate.After(nextBillingDate) || currentDate.Equal(nextBillingDate)
}

func CreateInvoiceLog(contract models.Contract, sendToCustomerStatus bool) {
	invoice := &models.Invoice{
		CustomerId:       contract.CustomerId,
		SubscriptionId:   contract.SubscriptionId,
		IssueDate:        time.Now(),
		Amount:           contract.InstallmentAmount,
		IsSendToCustomer: sendToCustomerStatus,
	}
	_, err := json.Marshal(invoice)
	if err != nil {
		fmt.Println("error occurred while marshalling invoice data: ", err)
	}
	value := dbCon.Create(invoice)
	if value.Error != nil {
		value.Rollback()
	} else {
		value.Commit()
	}

}

func DuplicateInvoiceSendCheck(contract models.Contract, billingDate time.Time) bool {
	newDbCon := db.GetDBConnection()
	invoice := []map[string]any{}
	firstDateOfCurMonth := time.Date(billingDate.Year(), billingDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(firstDateOfCurMonth, " ", contract.CustomerId, " ", contract.SubscriptionId, " ", billingDate)
	value := newDbCon.Model(&models.Invoice{}).Where("customer_id=? and subscription_id=? and issue_date>=? and issue_date<=? and is_send_to_customer=true", contract.CustomerId, contract.SubscriptionId, firstDateOfCurMonth, billingDate).Scan(&invoice)
	if value.Error != nil {
		fmt.Println("error occurred while fetching invoice data: ", value.Error.Error())
		return false
	}
	if len(invoice) > 0 {
		return true
	}
	return false
}

func GetActiveCustomers() {
	var contracts []models.Contract
	dbConE := db.GetDBConnection()
	value := dbConE.Preload("Customer").Preload("Subscription").Where("deleted_at is NULL and status=1").Find(&contracts)
	if value.Error != nil {
		fmt.Println(value.Error.Error())
	}
	for _, contract := range contracts {
		customerIdStr := strconv.Itoa(contract.CustomerId)
		nextBillingDate := getNextBillingDate(contract.ContractStartDate.UTC())
		if isBillingDue(nextBillingDate) && !DuplicateInvoiceSendCheck(contract, time.Now()) {
			err := generate_pdf.GeneratePDF(contract, "contract_details_"+customerIdStr+".pdf")
			if err != nil {
				fmt.Println("Error generating pdf: ", err)
				CreateInvoiceLog(contract, false)
			}
			CreateInvoiceLog(contract, true)
		}

	}

}

func main() {
	c := cron.New()
	errCron := c.AddFunc("0 12 * * *", GetActiveCustomers)
	if errCron != nil {
		fmt.Println("Error adding cron job:", errCron)
		return
	}
	c.Start()
	select {}
}
