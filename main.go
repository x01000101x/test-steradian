package main

import (
	"github.com/gin-gonic/gin"

	"crud-cars/controllers"
	"crud-cars/database"
	"crud-cars/models"
)

func main() {
	// Connect DB
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Car{}, &models.Order{})

	//Router
	r := gin.Default()

	//CARS
	//GET
	r.GET("/cars", controllers.GetCars)
	//CREATE
	r.POST("/cars", controllers.CreateCars)
	//UPDATE
	r.PUT("/cars/:car_id", controllers.UpdateCars)
	//DELETE
	r.DELETE("/cars/:car_id", controllers.DeleteCars)

	//ORDERS
	//GET
	r.GET("/orders", controllers.GetOrders)
	//CREATE
	r.POST("/orders", controllers.CreateOrders)
	//UPDATE
	r.PUT("/orders/:order_id", controllers.UpdateOrders)
	//DELETE
	r.DELETE("/orders/:order_id", controllers.DeleteOrders)

	//Server Start
	r.Run(":8080")

}
