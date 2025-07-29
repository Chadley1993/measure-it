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
		gpsResponse := services.GetGPSSamplingRate()
		context.JSON(http.StatusOK, gpsResponse)
		return
	}
	context.Status(http.StatusNoContent)
}
