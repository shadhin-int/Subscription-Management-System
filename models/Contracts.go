package models

import (
	"gorm.io/gorm"
	"time"
)

type Interval string
type Status string

const (
	IntervalMonths Interval = "month"
	IntervalYears  Interval = "yearly"
)

const (
	StatusActive    Status = "Active"
	StatusExpired   Status = "Expired"
	StatusCancelled Status = "Cancelled"
)

type Contract struct {
	gorm.Model
	ID                uint `gorm:"primaryKey" json:"id"`
	CustomerId        int
	SubscriptionId    int
	Customer          Customer     `gorm:"foreignKey:CustomerId" json:"customer"`
	Subscription      Subscription `gorm:"foreignKey:SubscriptionId" json:"subscription"`
	BillingInterval   Interval     `gorm:"BillingInterval" json:"billing_interval"`
	Status            Status       `gorm:"Status" json:"status"`
	ContractStartDate time.Time    `gorm:"contractStartDate" json:"contract_start_date"`
	ContractEndDate   time.Time    `gorm:"contractEndDate" json:"contract_end_date"`
	CreatedAt         time.Time    `gorm:"createdAt" json:"created_at"`
	UpdatedAt         time.Time    `gorm:"updatedAt" json:"updated_at"`
}
