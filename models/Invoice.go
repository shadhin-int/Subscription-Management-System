package models

import "time"

type InvoiceStatus string

const (
	InvoiceStatusActive    InvoiceStatus = "Issued"
	InvoiceStatusExpired   InvoiceStatus = "Paid"
	InvoiceStatusCancelled InvoiceStatus = "Overdue"
)

type Invoice struct {
	Id             uint          `gorm:"primaryKey" json:"id"`
	CustomerId     Customer      `gorm:"foreignKey:CustomerId" json:"customer"`
	SubscriptionId Subscription  `gorm:"foreignKey:SubscriptionId" json:"subscription"`
	IssueDate      time.Time     `gorm:"IssueDate" json:"issueDate"`
	DueDate        time.Time     `gorm:"DueDate" json:"due_date"`
	Amount         float64       `gorm:"Amount" json:"amount"`
	Status         InvoiceStatus `gorm:"InvoiceStatus" json:"status"`
}
