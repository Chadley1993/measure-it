package models

type ConfigData struct {
	GPSSampleRate int32  `json:"gpsSampleRate"`
	RPiHostIP     string `json:"rpiHostIP"`
	ActiveRecord  bool   `json:"activeRecord"`
}
