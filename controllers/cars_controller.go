package controllers

import (
	"crud-cars/database"
	"crud-cars/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get
func GetCars(c *gin.Context) {
	var cars []models.Car
	database.DB.Raw("SELECT * FROM cars").Find(&cars)
	c.JSON(http.StatusOK, cars)
}

// Creates
func CreateCars(c *gin.Context) {
	var cars models.Car
	if err := c.ShouldBindJSON(&cars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&cars)
	c.JSON(http.StatusOK, cars)
}

// Update
func UpdateCars(c *gin.Context) {
	car_id := c.Param("car_id")
	var cars models.Car
	if err := database.DB.Raw("SELECT * FROM cars WHERE id = ?", car_id).First(&cars, car_id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cars Not Found"})
		return
	}
	if err := c.ShouldBindJSON(&cars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&cars)
	c.JSON(http.StatusOK, cars)
}

// Delete
func DeleteCars(c *gin.Context) {
	id := c.Param("car_id")
	var cars models.Car
	if err := database.DB.Raw("SELECT * FROM cars WHERE id = ?", id).First(&cars, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cars Not Found"})
		return
	}
	database.DB.Delete(&cars)
	c.JSON(http.StatusOK, gin.H{"message": "Success Delete"})
}
