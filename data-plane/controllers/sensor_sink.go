package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/models"
	"measure-it.com/central-server/services"
)

func PostSensorData(context *gin.Context) {
	var data models.Dataframe
	if err := context.ShouldBindJSON(&data); err != nil {
		fmt.Println(context.Request.Body)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "errorMessage": err.Error()})
		return
	}
	services.UpdateSensorData(data)
	context.Status(http.StatusOK)
}
