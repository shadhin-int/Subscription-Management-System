package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"Name" json:"name"`
	Email     string    `gorm:"email;unique" json:"email"`
	CreatedAt time.Time `gorm:"createdAt" json:"created_at"`
	UpdatedAt time.Time `gorm:"updatedAt" json:"updated_at"`
}
