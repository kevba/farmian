package models

import "github.com/google/uuid"

type Farm struct {
	ID     uint32 `json:"id"`
	UserID uint32 `json:"userId"`

	Name     string `json:"name"`
	Location [2]int `json:"location"`

	WheatStock  int `json:"wheat"`
	PotatoStock int `json:"potato"`
	CornStock   int `json:"corn"`

	BarnLevel int `json:"barnLevel"`
}

func NewBarn(name string, location [2]int, userId uint32) *Farm {
	return &Farm{
		ID:          uuid.New().ID(),
		UserID:      userId,
		Name:        name,
		Location:    location,
		WheatStock:  0,
		PotatoStock: 0,
		CornStock:   0,
		BarnLevel:   0,
	}
}
