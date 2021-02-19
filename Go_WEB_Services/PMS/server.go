package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Mini Project": "Park in And Go",
		"Database":     "MySQL",
		"Framework":    "Gin Gonic",
	})
}

func RequestHandler() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/registration", controller.UserRegistration)
	r.POST("/login", controller.UserLogin)
	r.POST("/parkin", controller.VehicleRegistration)
	r.GET("/vehicle/:vehiclenumber", controller.VehicleDetails)
	r.POST("/vehicle/:tokennumber", controller.VehicleParkOut)
	r.GET("/all", controller.AllVehicle)
	r.Run(":8080")
}

func main() {
	RequestHandler()
}
