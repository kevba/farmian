package models

import "github.com/google/uuid"

type Farm struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`

	Name     string `json:"name"`
	Location [2]int `json:"location"`

	WheatStock  int `json:"wheat"`
	PotatoStock int `json:"potato"`
	CornStock   int `json:"corn"`

	Modifiers []string `json:"modifiers"`

	// Modifier controlled
	MaxStorage int `json:"maxStorage"`
}

func NewFarm(name string, location [2]int, userId string) *Farm {
	return &Farm{
		ID:          uuid.New().String(),
		UserID:      userId,
		Name:        name,
		Location:    location,
		WheatStock:  0,
		PotatoStock: 0,
		CornStock:   0,
		MaxStorage:  0,
		Modifiers:   []string{farmModifiers["BARN_LEVEL_0"].ID()},
	}
}
