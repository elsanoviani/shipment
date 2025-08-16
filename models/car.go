package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	ID     uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Brand  string    `gorm:"not null" json:"brand"`
	Model  string    `gorm:"not null" json:"model"`
	Year   int       `gorm:"not null" json:"year"`
	Price  float64   `gorm:"not null" json:"price"`
	Status string    `gorm:"default:'available'" json:"status"`
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
