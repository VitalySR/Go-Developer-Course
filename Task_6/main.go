package main

import (
	"fmt"
)

func main() {
	task1()

	task2()
}

type contract struct {
	ID     int
	Number string
	Date   string
}

func task1() {
	ct := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v\n", ct)

	ct1 := contract{
		ID:     1,
		Number: "#000A\n101",
		Date:   "2024-01-31",
	}
	fmtct1 := fmt.Sprintf("{ID:%d Number:%q Date:%s}", ct1.ID, ct1.Number, ct1.Date)
	fmt.Println(fmtct1)
}

func task2() {
	type contract struct {
		ID     int
		Number string
		Date   string
	}

	ct := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v\n", ct)
}
