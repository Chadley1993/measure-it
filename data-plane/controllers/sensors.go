package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/models"
	"measure-it.com/central-server/services"
)

func PostSensorData(context *gin.Context) {
	var data models.Dataframe
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}
	services.UpdateSensorData(data)
	gpsPoint := services.Point{
		X: (data.Latitude - 33.82) * 75000,
		Y: (data.Longitude - 18.526) * 100000,
	}
	services.ProcessGPSData(gpsPoint)
	context.Status(http.StatusNoContent)
}

func GetCarData(c *gin.Context) {
	sensorName := c.QueryArray("sensorName")
	c.PureJSON(http.StatusOK, services.GetRequestedData(sensorName))
}

func GetAllData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.PureJSON(http.StatusOK, services.GetSensorStore())
}
