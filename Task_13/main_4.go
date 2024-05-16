package main

import (
	"encoding/xml"
	"fmt"
)

type contracts struct {
	List []contract `xml:"contract"`
}

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}
	cs := contracts{}
	cs.List = append(cs.List, c)

	st, err := xml.MarshalIndent(cs, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(xml.Header, string(st))
}
