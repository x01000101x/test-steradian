package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId         uint   `json:"order_id"`
	CarId           int    `json:"car_id"`
	OrderDate       string `json:"order_date"`
	PickupDate      string `json:"pickup_date"`
	DropoffDate     string `json:"dropoff_date"`
	PickupLocation  string `json:"pickup_location"`
	DropoffLocation string `json:"dropoff_location"`
}
