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
	fmt.Println("-- Task5")
	task5()
}

func task1() {
	var m = map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Println(m)
}

func task2() {
	var m = make(map[string]int, 4)
	m["слон"] = 3
	m["бегемот"] = 0
	m["носорог"] = 5
	m["лев"] = 1
	fmt.Println(m)

	s := [3]string{"слон", "бегемот", "осьминог"}
	animal := s[0]
	cnt, ok := m[animal]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", animal, cnt, ok)
	animal = s[1]
	cnt, ok = m[animal]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", animal, cnt, ok)
	animal = s[2]
	cnt, ok = m[animal]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", animal, cnt, ok)
}

func task3() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Println(m)
	delete(m, "бегемот")
	fmt.Println(m)
}

func task4() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Println(m)
	m["выдра"] = struct{}{}
	fmt.Println(m)
}

func task5() {
	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	fmt.Println(m)
	m["бегемот"] = 2
	fmt.Println(m)
}
