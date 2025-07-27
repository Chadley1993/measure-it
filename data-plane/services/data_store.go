package services

import (
	"fmt"
	"sync"
	"time"

	"measure-it.com/central-server/models"
)

var (
	instance map[string]models.Dataframe
	once     sync.Once
)

func GetSensorStore() map[string]models.Dataframe {
	once.Do(func() {
		fmt.Println("Create map!!!")
		instance = make(map[string]models.Dataframe)
	})
	return instance
}

func GetRequestedData(requestedSensorNames []string) map[string]models.Dataframe {
	response := make(map[string]models.Dataframe)

	if len(requestedSensorNames) == 0 {
		return GetSensorStore()
	}

	sensorStore := GetSensorStore()
	for _, sensorName := range requestedSensorNames {
		if data, ok := sensorStore[sensorName]; ok {
			response[sensorName] = data
		}
	}
	return response
}

func UpdateSensorData(data models.Dataframe) {
	dataStore := GetSensorStore()
	data.Timestamp = time.Now()
	dataStore[data.SensorName] = data
}
