package controllers

import (
	"net/http"
	"strings"

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

	if strings.HasPrefix(data.SensorName, "gps-") {
		gpsResponse := services.GetConfig()
		context.JSON(http.StatusOK, gpsResponse)
		return
	}
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
