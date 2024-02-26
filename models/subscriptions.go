package models

import (
	"gorm.io/gorm"
	"time"
)

type DurationUnit string

const (
	DurationUnitMonths DurationUnit = "month"
	DurationUnitYears  DurationUnit = "year"
)

type Subscription struct {
	gorm.Model
	Id           uint         `gorm:"primaryKey" json:"id"`
	Name         string       `gorm:"Name" json:"name"`
	Code         string       `gorm:"code;index;unique" json:"code"`
	Description  string       `gorm:"Description" json:"description"`
	Price        float64      `gorm:"Price" json:"price"`
	Duration     int          `gorm:"Duration" json:"duration"`
	DurationUnit DurationUnit `gorm:"DurationUnit" json:"duration_unit"`
	IsActive     bool         `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time    `gorm:"createdAt" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"updatedAt" json:"updated_at"`
}
