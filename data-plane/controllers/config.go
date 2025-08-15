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

// func PostActiveRecording(c *gin.Context) {
// 	var data models.ConfigData
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
// 		return
// 	}
// 	services.ConfigMap.Store("activeRecord", data.ActiveRecord)
// 	c.Status(http.StatusOK)
// }

// func GetActiveRecording(c *gin.Context) {
// 	data, err := services.GetActiveRecording()
// 	if err != nil {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no host ip available", "errorMessage": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, data)
// }
