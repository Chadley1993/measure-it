package services

import (
	"fmt"

	"measure-it.com/central-server/models"
)

func GetConfigStore() *models.ConfigData {
	once.Do(func() {
		fmt.Println("Create ConfigStore map!!!")
		ConfigStore = models.ConfigData{
			// use as defaults
			GPSSampleRate: 20,
			RPiHostIP:     "x.x.x.x",
			ActiveRecord:  false,
		}

	})
	return &ConfigStore
}

func GetConfig() (sampleRate *models.ConfigData) {
	gpsResponse := GetConfigStore()
	return gpsResponse
}

func UpdateSamplingRate(newSamplingRate int32) {
	gpsResponse := GetConfigStore()
	gpsResponse.GPSSampleRate = newSamplingRate
}

func UpdateRPiHostIP(rpiIP string) {
	gpsResponse := GetConfigStore()
	gpsResponse.RPiHostIP = rpiIP
}

// This more STATE than it is Config to be fair
