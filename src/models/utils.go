package models

import (
	"encoding/json"
	"fmt"
	"os"
)

var saveDatapath = "./test_save_data"

func GetUsers() []*User {
	userData, _ := os.ReadFile(fmt.Sprintf("%s/accounts.json", saveDatapath))

	var users []*User

	_ = json.Unmarshal(userData, &users)

	return users
}
