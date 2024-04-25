package main

import (
	"fmt"
)

func main() {
	// fmt.Println("-- Task1")
	// task1()
	// fmt.Println("-- Task2")
	// task2()
	fmt.Println("-- Task3")
	task3()
}

func task1() {
	var fruit string
	fmt.Println("Введите наименование фрукта")
	fmt.Scanln(&fruit)
	cnt, ok := fruitMarket(fruit)
	if ok {
		fmt.Printf("%s: количество = %d\n", fruit, cnt)
	} else {
		fmt.Printf("%s: таких фруктов нет\n", fruit)
	}

}

func fruitMarket(fruit string) (int, bool) {
	var fruitMap = map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}
	cnt, ok := fruitMap[fruit]
	return cnt, ok
}

func task2() {
	num := []int{1, 2, 3}
	// fmt.Println(num)
	for a, ln := 0, len(num); a < ln; a++ {
		fmt.Printf("v1:%d\n", num[a])
	OUT:
		for b := 0; b < ln; b++ {
			fmt.Printf("\tv2:%d\n", num[b])
			for c := 0; c < ln; c++ {
				fmt.Printf("\t\tv3:%d\n", num[c])
				for d := 0; d < ln; d++ {
					fmt.Printf("\t\t\tv4:%d\n", num[d])
					if d >= 1 {
						break OUT
					}
				}
			}
		}
	}
}

func task3() {
	var fruit string
	fmt.Println("Введите наименование фрукта")
	fmt.Scanln(&fruit)
	fmt.Print(fruit + " - ")
	checkFood(fruit)
}

func checkFood(fruit string) {
	var fruitType uint8 = 0
	switch fruit {
	case "груша":
		fruitType = 1
	case "яблоко":
		fruitType = 1
	case "апельсин":
		fruitType = 1
	case "тыква":
		fruitType = 2
	case "огурец":
		fruitType = 2
	case "помидор":
		fruitType = 2
	default:
		fmt.Println("что-то странное...")
		return
	}

	if fruitType == 1 {
		fmt.Println("это фрукт")
		return
	}

	if fruitType == 2 {
		fmt.Println("это овощь")
	}
}
