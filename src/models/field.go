package models

type Field struct {
	ID       uint32 `json:"id"`
	FarmID   uint32 `json:"farmID"`
	Location [2]int `json:"location"`

	Crop     string `json:"crop"`
	SoilType string `json:"soilType"`
}
