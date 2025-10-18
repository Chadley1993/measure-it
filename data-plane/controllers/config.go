package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/models"
	"measure-it.com/central-server/services"
)

var activeRecordStatus chan struct{}
var currentStatus bool

func PatchGPSSampleRate(c *gin.Context) {
	var data models.ConfigData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}
	services.UpdateSamplingRate(data.GPSSampleRate)
	c.Status(http.StatusOK)
}

func PostRPiHostIP(c *gin.Context) {
	var data models.ConfigData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}
	services.UpdateRPiHostIP(data.RPiHostIP)
	c.Status(http.StatusOK)
}

func GetConfig(c *gin.Context) {
	data := services.GetConfigStore()
	c.JSON(http.StatusOK, data)
}

func PostDataRecording(context *gin.Context) {
	var data models.ConfigData
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}

	if !currentStatus && data.ActiveRecord {
		activeRecordStatus = make(chan struct{})
		go services.PushDataToNoSql(&activeRecordStatus)
		currentStatus = true
	} else {
		if currentStatus {
			close(activeRecordStatus)
		}
		currentStatus = false
	}

	context.Status(http.StatusOK)
}

// Just using > 20 sampling rate for now

func Test(context *gin.Context) {
	// var position services.Point
	// if err := context.ShouldBindJSON(&position); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
	// 	return
	// }
	var data models.Dataframe
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}

	if data.Latitude == 0 && data.Longitude == 0 {
		context.Status(http.StatusOK)
		return
	}

	gpsPoint := services.Point{
		Y: (data.Latitude - 33.82) * 75000,
		X: (data.Longitude - 18.526) * 100000,
	}
	services.ProcessGPSData(gpsPoint)
	context.Status(http.StatusOK)
}

func StartSession(context *gin.Context) {
	services.StartUp()
	context.Status(http.StatusOK)
}

func Health(context *gin.Context) {
	context.Status(http.StatusOK)
}
