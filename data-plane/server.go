package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/controllers"
	"measure-it.com/central-server/services"
)

func main() {
	fmt.Println("Starting server...")
	services.StartUp()
	router := gin.Default()

	router.GET("/sensorData", controllers.GetCarData)
	router.GET("/allSensorData", controllers.GetAllData)
	router.POST("/sensorData", controllers.PostSensorData)

	router.POST("/dataRecording", controllers.PostDataRecording)
	router.POST("/startSession", controllers.StartSession)
	router.POST("/test", controllers.Test)
	router.POST("/health", controllers.Health)

	router.GET("/config", controllers.GetConfig)
	router.PATCH("/config/rpi-ip", controllers.PostRPiHostIP)
	router.PATCH("/config/gps-samapling-rate", controllers.PatchGPSSampleRate)

	router.POST("/simulator", controllers.PostStartSimulator)
	router.PATCH("/simulator", controllers.PatchSimulator)
	router.DELETE("/simulator", controllers.DeleteSimulator)

	router.Run("0.0.0.0:8080")
}
