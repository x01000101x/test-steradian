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

	//GET
	r.GET("/cars", controllers.GetCars)
	//CREATE
	r.POST("/cars", controllers.CreateCars)
	//UPDATE
	r.PUT("/cars/:car_id", controllers.UpdateCars)
	//DELETE
	r.DELETE("/cars/:car_id", controllers.DeleteCars)

	//Server Start
	r.Run(":8080")

}
