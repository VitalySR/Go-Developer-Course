package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlor"`
	tenat    string
}

func main() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура",
	}

	s, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))
}
