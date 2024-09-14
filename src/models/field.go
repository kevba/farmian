package models

type Field struct {
	ID       string `json:"id"`
	FarmID   string `json:"farmID"`
	Location [2]int `json:"location"`

	Crop     string `json:"crop"`
	SoilType string `json:"soilType"`
}
