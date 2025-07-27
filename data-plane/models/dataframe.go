package models

import "time"

type Dataframe struct {
	SensorName string    `json:"sensorName" binding:"required"`
	Timestamp  time.Time `json:"tmstamp,omitempty"`

	SpeedKPH    float64 `json:"speedKPH,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Temperature string  `json:"temperature,omitempty"`
}
