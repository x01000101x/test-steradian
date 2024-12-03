package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID        uint    `gorm:"car_id"`
	CarName   string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}
