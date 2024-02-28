package models

import "time"

type Invoice struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	CustomerId       int       `gorm:"customerId;index" json:"customer_id"`
	SubscriptionId   int       `gorm:"subscriptionId;index" json:"subscription_id"`
	IssueDate        time.Time `gorm:"issueDate" json:"issue_date"`
	DueDate          time.Time `gorm:"dueDate" json:"due_date"`
	Amount           float64   `gorm:"amount" json:"amount"`
	IsSendToCustomer bool      `gorm:"IsSendToCustomer" json:"is_send_to_customer"`
}
