package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	db "subscription_management_system/database"
	"time"
)

type Contract struct {
	ID                uint
	CustomerID        int
	SubscriptionID    int
	BillingInterval   int8
	InstallmentAmount float64
	Duration          int
	DurationUnit      int8
	Status            int8
	ContractStartDate time.Time
	ContractEndDate   time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func Customer() chi.Router {
	route := chi.NewRouter()

	return route
}

func Invoice() chi.Router {
	route := chi.NewRouter()

	return route
}

func getNextBillingDate(contractStartDate time.Time) time.Time {
	// Get the year and month of the contract start date
	year, month, _ := contractStartDate.Date()

	// Get the next month
	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		year++
	}

	// Get the number of days in the next month
	nextMonthDays := time.Date(year, nextMonth, 0, 0, 0, 0, 0, time.UTC).Day()

	// Create the next billing date by using the same day as the contract start date
	nextBillingDate := time.Date(year, nextMonth, contractStartDate.Day(), 0, 0, 0, 0, time.UTC)

	// If the contract start day is greater than the number of days in the next month,
	// adjust the billing date to the last day of the next month
	if contractStartDate.Day() > nextMonthDays {
		nextBillingDate = time.Date(year, nextMonth, nextMonthDays, 0, 0, 0, 0, time.UTC)
	}

	return nextBillingDate
}

func isBillingDue(nextBillingDate time.Time) bool {
	// Get the current date
	currentDate := time.Now()
	fmt.Println(nextBillingDate, " ", currentDate)
	// Check if the current date is greater than or equal to the next billing date
	return currentDate.After(nextBillingDate) || currentDate.Equal(nextBillingDate)
}

func GetActiveCustomers() []Contract {
	var contracts []Contract
	var dueContractList []Contract
	dbCon := db.GetDBConnection()
	value := dbCon.Where("deleted_at is NULL and status=1").Find(&contracts)
	if value.Error != nil {
		fmt.Println(value.Error.Error())
	}
	fmt.Println(contracts)
	for _, contract := range contracts {
		//parsedTime, err := time.Parse(time.RFC3339, contract.ContractStartDate)
		//if err != nil {
		//	fmt.Println("Error parsing date string:", err)
		//	return nil
		//}
		loc, err := time.LoadLocation("Asia/Dhaka")
		if err != nil {
			fmt.Println("Error loading location:", err)
			return nil
		}
		localTime := contract.ContractStartDate.In(loc)
		formatted := localTime.Format("2006-01-02 15:04:05 MST")
		fmt.Println(formatted)

		nextBillingDate := getNextBillingDate(localTime)
		//fmt.Println("data: ", parsedTime, " ", nextBillingDate)

		if isBillingDue(nextBillingDate) {
			dueContractList = append(dueContractList, contract)
		}
	}
	return dueContractList

}

func main() {
	router := chi.NewRouter()
	okData := GetActiveCustomers()
	fmt.Println(okData)

	router.Mount("/customer/", Customer())
	router.Mount("/invoice/", Invoice())

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
