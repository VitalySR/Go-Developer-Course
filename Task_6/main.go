package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Task1")
	task1()
	fmt.Println("-- Task2")
	task2()
	fmt.Println("-- Task3")
	task3()
	fmt.Println("-- Task4")
	task4()
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

func task3() {
	cont := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}
	fmt.Println(cont.print())
}

func (ct contract) print() string {
	return fmt.Sprintf("Договор № %s от %s", ct.Number, ct.Date)
}

func task4() {
	type contact struct {
		Addss string
		Phone string
	}
	type user struct {
		ID   int
		Name string
		contact
	}
	type employee struct {
		ID   int
		Name string
		contact
	}

	u := user{
		ID:   1,
		Name: "Коля",
		contact: contact{
			Addss: "г.Лысьва",
			Phone: "+7(950)222-77-88",
		},
	}

	e := employee{
		ID:   1,
		Name: "Петя",
		contact: contact{
			Addss: "г.Пермь",
			Phone: "+7(909)111-33-44",
		},
	}
	// fmt.Printf("%+v\n", u)
	// fmt.Printf("%+v\n", e)
	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
