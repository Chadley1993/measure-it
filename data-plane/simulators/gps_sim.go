package simulators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"measure-it.com/central-server/models"
)

type GPSSimConfig struct {
	Min         int
	Max         int
	FrequenceMS int
	RandomNoise int
}

func GPSSimulator(name string, stop chan int8, update chan GPSSimConfig) {
	simConfig := GPSSimConfig{
		Min:         60,
		Max:         180,
		FrequenceMS: 2000,
		RandomNoise: 0,
	}

	gpsData := models.Dataframe{
		SensorName: name,
	}

	for {
		select {
		case simConfig = <-update:
			fmt.Println("Update ->", simConfig)
		case <-stop:
			return
		default:
			gpsData.SpeedKPH = rand.Int63() * 100
			jsonData, err := json.Marshal(gpsData)
			if err != nil {
				fmt.Println("Error Marshalling data:", err)
				return
			}

			url := "http://localhost:8080/sensorData"
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}

			if resp.StatusCode != 200 {
				fmt.Println("Failed request, status code:", resp.StatusCode)
				return
			}

			defer resp.Body.Close()
			time.Sleep(time.Millisecond * time.Duration(simConfig.FrequenceMS))
		}
	}
}
