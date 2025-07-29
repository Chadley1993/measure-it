package services

import (
	"fmt"

	"measure-it.com/central-server/models"
)

func GetGPSSamplingRate() (sampleRate models.GPSFeedback) {
	gpsResponse := models.GPSFeedback{}
	if val, ok := ConfigMap.Load("gpsSamplingRate"); ok {
		gpsResponse.SampleRate = val.(float32)
		return gpsResponse
	}
	fmt.Println("Using default sampling rate")
	gpsResponse.SampleRate = 20 //Default value
	return gpsResponse
}
