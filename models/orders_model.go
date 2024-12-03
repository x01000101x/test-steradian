package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId         int    `json:"orderId"`
	CarId           string `json:"carId"`
	OrderDate       int    `json:"orderDate"`
	PickupDate      int    `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
}
