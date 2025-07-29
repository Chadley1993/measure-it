package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"measure-it.com/central-server/services"
)

func GetCarData(c *gin.Context) {
	sensorName := c.QueryArray("sensorName")
	fmt.Println("Params:", sensorName)
	c.PureJSON(http.StatusOK, services.GetRequestedData(sensorName))
}
