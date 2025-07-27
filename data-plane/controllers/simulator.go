package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/simulators"
)

var stopSignal = make(chan int8)
var simConfig = make(chan simulators.GPSSimConfig)
var isRunning bool

func PostStartSimulator(context *gin.Context) {
	sensorName := context.Query("sensorName")
	if sensorName == "" {
		fmt.Println("No sensor name.")
		context.Status(http.StatusBadRequest)
	}

	if !isRunning {
		go simulators.GPSSimulator(sensorName, stopSignal, simConfig)
		isRunning = true
		context.Status(http.StatusOK)
		return
	}
	context.Status(http.StatusConflict)
}

func PatchSimulator(context *gin.Context) {
	var newConfig simulators.GPSSimConfig
	if err := context.ShouldBindJSON(&newConfig); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body for GPSSimConfig", "errorMessage": err.Error()})
		return
	}

	simConfig <- newConfig
}

func DeleteSimulator(context *gin.Context) {
	defer func() {
		isRunning = false
	}()
	stopSignal <- 0
}
