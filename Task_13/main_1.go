package main

import (
	"encoding/json"
	"fmt"
)

type contact struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlor"`
	Tenat    string `json:"tenat"`
}

func main() {
	s := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	t := contact{}
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
