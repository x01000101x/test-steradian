package controllers

import (
	"crud-cars/database"
	"crud-cars/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get
func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Raw("SELECT * FROM orders").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// Creates
func CreateOrders(c *gin.Context) {
	var orders models.Order
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&orders)
	c.JSON(http.StatusOK, orders)
}

// Update
func UpdateOrders(c *gin.Context) {
	order_id := c.Param("order_id")
	var orders models.Order
	if err := database.DB.Raw("SELECT * FROM orders WHERE order_id = ?", order_id).First(&orders, order_id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orders Not Found"})
		return
	}
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&orders)
	c.JSON(http.StatusOK, orders)
}

// Delete
func DeleteOrders(c *gin.Context) {
	id := c.Param("order_id")
	var orders models.Order
	if err := database.DB.Raw("SELECT * FROM orders WHERE order_id = ?", id).First(&orders, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orders Not Found"})
		return
	}
	database.DB.Delete(&orders)
	c.JSON(http.StatusOK, gin.H{"message": "Success Delete"})
}
