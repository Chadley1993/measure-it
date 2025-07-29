package models

type GPSFeedback struct {
	SampleRate float32 `json:"sampleRate" binding:"required"`
}
