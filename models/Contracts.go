package models

import (
	"gorm.io/gorm"
	"time"
)

type Contract struct {
	gorm.Model
	ID                uint `gorm:"primaryKey" json:"id"`
	CustomerId        int
	SubscriptionId    int
	Customer          Customer     `gorm:"foreignKey:CustomerId" json:"customer"`
	Subscription      Subscription `gorm:"foreignKey:SubscriptionId" json:"subscription"`
	BillingInterval   int8         `gorm:"billingInterval" json:"billing_interval"`
	InstallmentAmount float64      `gorm:"installmentAmount" json:"installment_amount"`
	Duration          int          `gorm:"duration" json:"duration"`
	DurationUnit      int8         `gorm:"durationUnit" json:"duration_unit"`
	Status            int8         `gorm:"status" json:"status"`
	ContractStartDate time.Time    `gorm:"contractStartDate" json:"contract_start_date"`
	ContractEndDate   time.Time    `gorm:"contractEndDate" json:"contract_end_date"`
	CreatedAt         time.Time    `gorm:"createdAt" json:"created_at"`
	UpdatedAt         time.Time    `gorm:"updatedAt" json:"updated_at"`
}
