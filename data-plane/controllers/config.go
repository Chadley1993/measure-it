package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/models"
	"measure-it.com/central-server/services"
)

func PatchGPSSamplingRate(c *gin.Context) {
	var data models.GPSFeedback
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}
	services.ConfigMap.Store("gpsSamplingRate", data.SampleRate)
	c.Status(http.StatusOK)
}
